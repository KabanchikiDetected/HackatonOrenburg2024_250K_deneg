package server

import (
	"errors"
	"net/http"

	"github.com/KabanchikiDetected/HackatonOrenburg2024_250K_deneg/internal/domain/requests"
	"github.com/KabanchikiDetected/HackatonOrenburg2024_250K_deneg/internal/service"
)

func (s *Server) createProduct(w http.ResponseWriter, r *http.Request) {
	var req requests.CreateProduct
	err := Decode(w, r, &req)
	if err != nil {
		return
	}

	err = s.service.CreateProduct(r.Context(), req)
	if err != nil {
		handleError(err, w)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (s *Server) getProducts(w http.ResponseWriter, r *http.Request) {
	products, err := s.service.Products(r.Context())
	if err != nil {
		handleError(err, w)
		return
	}

	err = Encode(w, r, products)
	if err != nil {
		return
	}
}

func (s *Server) getProduct(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	product, err := s.service.Product(r.Context(), id)
	if err != nil {
		handleError(err, w)
		return
	}
	err = Encode(w, r, product)
	if err != nil {
		return
	}

}

func (s *Server) buyProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("product_id")
	payload, err := FromContext(r.Context())
	if err != nil {
		return
	}
	err = s.service.Buy(r.Context(), payload.ID, productId)
	if err != nil {
		handleError(err, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func handleError(err error, w http.ResponseWriter) {
	if errors.Is(err, service.ErrInternalServer) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if errors.Is(err, service.ErrBadRequest) {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if errors.Is(err, service.ErrNotFound) {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}
