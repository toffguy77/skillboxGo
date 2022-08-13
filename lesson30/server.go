package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type User struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Friends *[]User `json:"friends,omitempty"`
}

type service struct {
	store map[string]*User
}

func main() {
	var wg sync.WaitGroup
	// storage
	srv := service{make(map[string]*User)}

	wg.Add(1)
	go func() {
		defer wg.Done()

		// start server
		r := chi.NewRouter()
		r.Use(middleware.Logger)
		r.Use(middleware.Recoverer)

		r.HandleFunc("/add", srv.CreateUser)
		r.HandleFunc("/get", srv.GetUser)
		r.HandleFunc("/getall", srv.GetAllUsers)
		r.HandleFunc("/update", srv.UpdateUser)
		r.HandleFunc("/delete", srv.DeleteUser)

		if err := http.ListenAndServe(":8080", r); err != nil {
			log.Fatal("ListenAndServe: " + err.Error())
		}
	}()
	wg.Wait()
}

func (srv *service) CreateUser(w http.ResponseWriter, r *http.Request) {
	if !checkHeaders(r) {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte("invalid headers")); err != nil {
			log.Printf("can't send response: %v\n", err)
		}
		return
	}

	if r.Method == "POST" {
		content, err := io.ReadAll(r.Body)
		log.Println(string(content))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			if _, err := w.Write([]byte(err.Error())); err != nil {
				log.Printf("can't send response: %v\n", err)
			}
			return
		}

		var u *User
		if err = json.Unmarshal(content, &u); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			if _, err := w.Write([]byte(err.Error())); err != nil {
				log.Printf("can't send response: %v\n", err)
			}
			return
		}

		srv.store[u.Name] = u

		w.WriteHeader(http.StatusCreated)
		if _, err := w.Write([]byte("user was created: " + u.Name)); err != nil {
			log.Printf("can't send response: %v\n", err)
		}
	}

	w.WriteHeader(http.StatusBadRequest)
}

func (srv *service) GetUser(w http.ResponseWriter, r *http.Request) {
	if !checkHeaders(r) {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte("invalid headers")); err != nil {
			log.Printf("can't send response: %v\n", err)
		}
		return
	}

	if r.Method == "GET" {
		content, err := io.ReadAll(r.Body)
		log.Println(string(content))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			if _, err := w.Write([]byte(err.Error())); err != nil {
				log.Printf("can't send response: %v\n", err)
			}
			return
		}

		u := r.FormValue("name")
		user, ok := srv.store[u]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			if _, err := w.Write([]byte(err.Error())); err != nil {
				log.Printf("can't send response: %v\n", err)
			}
			return
		}

		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(user.toString())); err != nil {
			log.Printf("can't send response: %v\n", err)
		}
	}

	w.WriteHeader(http.StatusBadRequest)
}

func (srv *service) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	if !checkHeaders(r) {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte("not a json data request")); err != nil {
			log.Printf("can't send response: %v\n", err)
		}
		return
	}

	if r.Method == "GET" {
		response := ""
		for _, user := range srv.store {
			response += user.toString() + "\n"
			log.Println(response)
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(response)); err != nil {
			log.Printf("can't send response: %v\n", err)
		}
	}
}

func (srv *service) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if !checkHeaders(r) {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte("invalid headers")); err != nil {
			log.Printf("can't send response: %v\n", err)
		}
		return
	}

	if r.Method == "DELETE" {
		content, err := io.ReadAll(r.Body)
		log.Println(string(content))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			if _, err := w.Write([]byte(err.Error())); err != nil {
				log.Printf("can't send response: %v\n", err)
			}
			return
		}

		u := r.FormValue("name")
		if _, ok := srv.store[u]; !ok {
			w.WriteHeader(http.StatusNotFound)
			if _, err := w.Write([]byte(err.Error())); err != nil {
				log.Printf("can't send response: %v\n", err)
			}
			return
		}

		delete(srv.store, u)
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte("successfully deleted: " + u)); err != nil {
			log.Printf("can't send response: %v\n", err)
		}
	}

	w.WriteHeader(http.StatusBadRequest)
}

func (srv *service) UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func checkHeaders(req *http.Request) bool {
	if req.Header.Get("Content-Type") == "application/json; charset=utf-8" {
		return true
	}
	return false
}

func (u User) toString() string {
	return fmt.Sprintf("name:%s, age: %d", u.Name, u.Age)
}
