package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KabanchikiDetected/HackatonOrenburg2024_250K_deneg/internal/config"
	"github.com/KabanchikiDetected/HackatonOrenburg2024_250K_deneg/internal/domain/requests"
	"github.com/KabanchikiDetected/HackatonOrenburg2024_250K_deneg/internal/domain/responses"

	// "github.com/KabanchikiDetected/HackatonOrenburg2024_250K_deneg/internal/domain/requests"
	"github.com/go-pkgz/routegroup"
)

type Service interface {
	CreateProduct(ctx context.Context, product requests.CreateProduct) error
	Product(ctx context.Context, id string) (responses.Product, error)
	Products(ctx context.Context) ([]responses.Product, error)
	Buy(ctx context.Context, userId, productId string) error
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
		Addr:    fmt.Sprintf("%s:%d", config.Config().Server.Host, config.Config().Server.Port),
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
	jwtMV := MiddlwareJWT(config.Config().JWT.PublicKey)
	s.mux.HandleFunc("/products", s.getProducts)
	s.mux.HandleFunc("/product", s.getProduct)
	s.mux.With(jwtMV).HandleFunc("/buy", s.buyProduct)
	s.mux.With(jwtMV).HandleFunc("/create", s.createProduct)
}

func (s *Server) Handler() http.Handler {
	return s.mux
}
