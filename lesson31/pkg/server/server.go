package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"io"
	"log"
	"net/http"
	"skillbox/pkg/models"
	"skillbox/pkg/storage/mem"
	"strconv"
)

var store *mem.MemStorage

type Server struct {
	Router *chi.Mux
	// Db, config can be added here
}

type UserResponse struct {
	Status   int         `json:"status"`
	Response models.User `json:"response,omitempty"`
	Error    error       `json:"error,omitempty"`
}

type UsersResponse struct {
	Status   int           `json:"status"`
	Response []models.User `json:"response,omitempty"`
	Error    error         `json:"error,omitempty"`
}

func init() {
	store = mem.NewStorage()
}

func CreateNewServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	return s
}

func (s *Server) MountHandlers() {
	r := s.Router
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.CleanPath)

	r.Route("/", func(r chi.Router) {
		r.Get("/", DefaultRoute)

		// sub-route 'users'
		r.Route("/users", func(r chi.Router) {
			r.With(paginate).Get("/", GetAllUsers)
			r.Post("/create", CreateUser)
			r.Post("/make_friend", MakeFriend)

			// sub-route by user's id
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", GetUser)
				r.Put("/", UpdateUser)
				r.Delete("/", DeleteUser)
			})
		})

		// sub-route 'friends'
		r.Route("/friends", func(r chi.Router) {
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", GetFriends)
			})
		})
	})
}

func DefaultRoute(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("hello, there!")); err != nil {
		log.Printf("can't send response: %v\n", err)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	content, err := io.ReadAll(r.Body)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, models.User{}, err)
		return
	}

	var u *models.User
	if err = json.Unmarshal(content, &u); err != nil {
		sendResponse(w, http.StatusInternalServerError, models.User{}, err)
		return
	}

	userID, err := store.Save(u)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, *u, err)
		log.Printf("error saving user to the storage: %v\n", err)
		return
	}
	u.ID = userID

	sendResponse(w, http.StatusCreated, *u, nil)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	userIDString := chi.URLParam(r, "id")
	userID, _ := strconv.Atoi(userIDString)
	user := store.Get(userID)
	if user == nil {
		error404 := errors.New("user " + userIDString + " not found")
		sendResponse(w, http.StatusNotFound, models.User{}, error404)
		return
	}

	sendResponse(w, http.StatusOK, *user, nil)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := store.AllUsers()

	if users == nil {
		error404 := errors.New("users not found")
		sendResponse(w, http.StatusNotFound, models.User{}, error404)
		return
	}

	sendResponseSlice(w, http.StatusOK, users, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userIDString := chi.URLParam(r, "id")
	userID, _ := strconv.Atoi(userIDString)
	err := store.Delete(userID)
	if err != nil {
		sendResponse(w, http.StatusNotFound, models.User{}, err)
		return
	}

	sendResponse(w, http.StatusOK, models.User{}, nil)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userIDString := chi.URLParam(r, "id")
	userID, _ := strconv.Atoi(userIDString)

	content, err := io.ReadAll(r.Body)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, models.User{}, err)
		return
	}

	var u *models.User
	if err = json.Unmarshal(content, &u); err != nil {
		sendResponse(w, http.StatusInternalServerError, models.User{}, err)
		return
	}

	user, err := store.Update(userID, u)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, *user, err)
		return
	}

	sendResponse(w, http.StatusOK, *user, nil)
}

func MakeFriend(w http.ResponseWriter, r *http.Request) {
	content, err := io.ReadAll(r.Body)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, models.User{}, err)
		return
	}

	var req models.FriendRequest
	err = json.Unmarshal(content, &req)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, models.User{}, err)
		return
	}

	source := store.Get(req.Source_id)
	target := store.Get(req.Target_id)
	if &source != nil && &target != nil {
		newSource, err := source.MakeFriend(target)
		if err != nil {
			log.Println(err)
			return
		}
		store.Update(source.ID, newSource)
		updatedSource := store.Get(source.ID)

		sendResponse(w, http.StatusOK, *updatedSource, nil)
	}
	errIncomplete := errors.New(fmt.Sprintf("request is not complete: source %v, target %v", source, target))
	sendResponse(w, http.StatusUnprocessableEntity, models.User{}, errIncomplete)
}

func GetFriends(w http.ResponseWriter, r *http.Request) {
	userIDString := chi.URLParam(r, "id")
	userID, _ := strconv.Atoi(userIDString)
	user := store.Get(userID)
	if user == nil {
		error404 := errors.New("user " + userIDString + " not found")
		sendResponse(w, http.StatusNotFound, models.User{}, error404)
		return
	}

	var friends []models.User
	for _, user := range user.Friends {
		friends = append(friends, *user)
	}

	sendResponseSlice(w, http.StatusOK, friends, nil)
}

func paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func sendResponse(w http.ResponseWriter, status int, user models.User, err error) {
	w.WriteHeader(status)
	response := UserResponse{
		Status:   status,
		Response: user,
		Error:    err,
	}
	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Printf("can't prepare response: %v\n", err)
		return
	}
	if _, err := w.Write(jsonData); err != nil {
		log.Printf("can't send response: %v\n", err)
		return
	}
}

func sendResponseSlice(w http.ResponseWriter, status int, users []models.User, err error) {
	w.WriteHeader(status)
	response := UsersResponse{
		Status:   status,
		Response: users,
		Error:    err,
	}
	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Printf("can't prepare response: %v\n", err)
		return
	}
	if _, err := w.Write(jsonData); err != nil {
		log.Printf("can't send response: %v\n", err)
		return
	}
}
