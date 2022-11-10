package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"log"
	"net/http"
	"skillbox/internal/models"
)

type BaseHandler struct {
	userRepo models.Storage
}

func NewBaseHandler(userRepo models.Storage) *BaseHandler {
	return &BaseHandler{
		userRepo: userRepo,
	}
}

func (h *BaseHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, _ := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
	user := h.userRepo.Get(id)
	if user == nil {
		error404 := errors.New(fmt.Sprintf("user %v not found", id))
		sendResponse(w, http.StatusNotFound, models.User{}, error404)
		return
	}
	sendResponse(w, http.StatusOK, *user, nil)
}

func (h *BaseHandler) Save(w http.ResponseWriter, r *http.Request) {
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

	user, err := h.userRepo.Save(u)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, *user, err)
		log.Printf("error saving user to the storage: %v\n", err)
		return
	}
	sendResponse(w, http.StatusCreated, *user, nil)
}

func (h *BaseHandler) AllUsers(w http.ResponseWriter, r *http.Request) {
	allUsers := h.userRepo.AllUsers()

	if allUsers == nil {
		error404 := errors.New("users not found")
		sendResponse(w, http.StatusNotFound, models.User{}, error404)
		return
	}

	sendResponseSlice(w, http.StatusOK, allUsers, nil)
}

func (h *BaseHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
	err := h.userRepo.Delete(id)
	if err != nil {
		sendResponse(w, http.StatusNotFound, models.User{}, err)
		return
	}

	sendResponse(w, http.StatusOK, models.User{}, nil)
}

func (h *BaseHandler) Update(w http.ResponseWriter, r *http.Request) {
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

	user, err := h.userRepo.Update(u)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, *user, err)
		return
	}

	sendResponse(w, http.StatusOK, *user, nil)
}

func (h *BaseHandler) MakeFriend(w http.ResponseWriter, r *http.Request) {
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

	source := h.userRepo.Get(req.Source_id)
	target := h.userRepo.Get(req.Target_id)
	if &source != nil && &target != nil {
		newSource, err := h.userRepo.MakeFriend(source, target)
		if err != nil {
			errIncomplete := errors.New(fmt.Sprintf("request is not complete: source %v, target %v", source, target))
			sendResponse(w, http.StatusUnprocessableEntity, models.User{}, errIncomplete)
			return
		}

		sendResponse(w, http.StatusOK, *newSource, nil)
		return
	}
	errIncomplete := errors.New(fmt.Sprintf("request is not complete: source %v, target %v", source, target))
	sendResponse(w, http.StatusUnprocessableEntity, models.User{}, errIncomplete)
}

func (h *BaseHandler) DeleteFriend(w http.ResponseWriter, r *http.Request) {
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

	source := h.userRepo.Get(req.Source_id)
	target := h.userRepo.Get(req.Target_id)
	if &source != nil && &target != nil {
		newSource, err := h.userRepo.DeleteFriend(source, target)
		if err != nil {
			errIncomplete := errors.New(fmt.Sprintf("request is not complete: source %v, target %v", source, target))
			sendResponse(w, http.StatusUnprocessableEntity, models.User{}, errIncomplete)
			return
		}

		sendResponse(w, http.StatusOK, *newSource, nil)
	}
	errIncomplete := errors.New(fmt.Sprintf("request is not complete: source %v, target %v", source, target))
	sendResponse(w, http.StatusUnprocessableEntity, models.User{}, errIncomplete)
}

func (h *BaseHandler) GetFriends(w http.ResponseWriter, r *http.Request) {
	id, _ := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
	friends, err := h.userRepo.GetFriends(id)
	if err != nil {
		error404 := errors.New(fmt.Sprintf("user %v not found", id))
		sendResponse(w, http.StatusNotFound, models.User{}, error404)
		return
	}
	sendResponseSlice(w, http.StatusOK, friends, nil)
}

func sendResponse(w http.ResponseWriter, status int, user models.User, err error) {
	w.WriteHeader(status)
	response := models.UserResponse{
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
	response := models.UsersResponse{
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
