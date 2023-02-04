package server

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"skillbox/internal/controllers"
	"skillbox/internal/storage"
)

type Server struct {
	Router *chi.Mux
	// Db, config can be added here
}

func CreateNewServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	return s
}

func (s *Server) MountHandlers(ctx *context.Context) {
	r := s.Router
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.CleanPath)

	userRepo := storage.NewUserRepo(ctx, "friends-mongo")
	h := controllers.NewBaseHandler(userRepo)

	r.Route("/", func(r chi.Router) {
		r.Get("/", DefaultRoute)

		// sub-route 'users'
		r.Route("/users", func(r chi.Router) {
			r.With(paginate).Get("/", h.AllUsers)
			r.Post("/create", h.Save)

			// sub-route by user's id
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", h.Get)
				r.Put("/", h.Update)
				r.Delete("/", h.Delete)
			})
		})

		// sub-route 'friends'
		r.Route("/friends", func(r chi.Router) {
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", h.GetFriends)
				r.Post("/make_friend", h.MakeFriend)
				r.Delete("/", h.DeleteFriend)
			})
		})
	})
}

func DefaultRoute(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("hello, there!")); err != nil {
		log.Printf("can't send response: %v\n", err)
	}
}

func paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
