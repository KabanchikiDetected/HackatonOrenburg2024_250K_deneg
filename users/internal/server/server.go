package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KabanchikiDetected/users/internal/config"
	"github.com/KabanchikiDetected/users/internal/domain/requests"
	"github.com/go-pkgz/routegroup"
)

type Service interface {
	Register(ctx context.Context, user requests.Register) error
	Login(ctx context.Context, user requests.Login) (token string, err error)
}

type Server struct {
	server  *http.Server
	mux     *routegroup.Bundle
	service Service
}

func New(service Service) *Server {
	httpMux := http.NewServeMux()
	mux := routegroup.New(httpMux)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Config().Server.Port),
		Handler: mux,
	}
	s := &Server{
		server:  server,
		mux:     mux,
		service: service,
	}
	s.initServer()
	return s
}

func (s *Server) Run() {
	fmt.Printf("Server started on %s\n", s.server.Addr)
	if err := s.server.ListenAndServe(); err != nil {
		fmt.Println("Error: ", err)
	}
}

func (s *Server) Stop() {
	if err := s.server.Shutdown(context.Background()); err != nil {
		fmt.Println("Error:", err)
	}
}

func (s *Server) initServer() {
	s.mux.Use(MiddlewareCors)
	s.mux.HandleFunc("POST /login", s.login)
	s.mux.HandleFunc("POST /register", s.register)
}

func (s *Server) Handler() http.Handler {
	return s.mux
}
