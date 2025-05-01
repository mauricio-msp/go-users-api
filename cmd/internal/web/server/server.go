package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mauricio-msp/go-users-api/cmd/internal/service"
	"github.com/mauricio-msp/go-users-api/cmd/internal/web/handlers"
)

type Server struct {
	router      *chi.Mux
	server      *http.Server
	userService *service.UserService
}

func NewServer(userService *service.UserService) *Server {
	return &Server{
		router:      chi.NewRouter(),
		userService: userService,
	}
}

func (s *Server) setupMiddlewares() {
	s.router.Use(middleware.Recoverer)
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.Logger)
}

func (s *Server) Routes() {
	userHandler := handlers.NewUserHandler(s.userService)

	s.setupMiddlewares()

	s.router.Route("/api", func(r chi.Router) {
		r.Post("/users", userHandler.CreateUser)
		r.Get("/users", userHandler.ListUsers)
		r.Get("/users/{id}", userHandler.GetUser)
		r.Delete("/users/{id}", userHandler.DeleteUser)
		r.Put("/users/{id}", userHandler.UpdateUser)
	})
}

func (s *Server) Start() error {
	s.server = &http.Server{
		Addr:         ":3333",
		Handler:      s.router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	return s.server.ListenAndServe()
}
