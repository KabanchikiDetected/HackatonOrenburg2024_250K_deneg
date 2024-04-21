package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KabanchikiDetected/hackaton/students/internal/config"
	"github.com/KabanchikiDetected/hackaton/students/internal/server/middlewares"
	"github.com/KabanchikiDetected/hackaton/students/internal/server/users"
	"github.com/go-pkgz/routegroup"
)

type Server struct {
	server *http.Server
	mux    *routegroup.Bundle
}

func New(
	usersService users.UsersService,
) *Server {
	cfg := config.Config()
	httpMx := http.NewServeMux()

	mux := routegroup.New(httpMx)
	mux.Use(middlewares.SetHeaders)
	mux.Use(middlewares.MiddlewareCors)

	users.Register(usersService, mux.Mount("/users"))

	return &Server{
		server: &http.Server{
			Addr:    fmt.Sprintf(":%s", cfg.Server.Port),
			Handler: mux,
		},
		mux: mux,
	}
}

func (s *Server) Start() {
	fmt.Printf("Starting server on %s\n", s.server.Addr)
	if err := s.server.ListenAndServe(); err != nil {
		fmt.Println("Error: ", err)
	}
}

func (s *Server) Stop() {
	if err := s.server.Shutdown(context.Background()); err != nil {
		fmt.Println("Error: ", err)
	}
}

func (s *Server) Handler() http.Handler {
	return s.mux
}
