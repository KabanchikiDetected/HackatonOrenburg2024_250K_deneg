package server

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/KabanchikiDetected/HackatonOrenburg2024_250K_deneg/users/internal/domain/requests"
	"github.com/KabanchikiDetected/HackatonOrenburg2024_250K_deneg/users/internal/domain/responses"
	"github.com/KabanchikiDetected/HackatonOrenburg2024_250K_deneg/users/internal/service"
	"github.com/Richtermnd/utilshttp"
)

func (s *Server) register(w http.ResponseWriter, r *http.Request) {
	var req requests.Register
	err := utilshttp.Decode(w, r, &req)
	if err != nil {
		return
	}

	if err := s.service.Register(r.Context(), req); err != nil {
		handleError(err, w)
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *Server) login(w http.ResponseWriter, r *http.Request) {
	var req requests.Login
	err := utilshttp.Decode(w, r, &req)
	if err != nil {
		return
	}

	token, err := s.service.Login(r.Context(), req)
	if err != nil {
		handleError(err, w)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responses.Token{Token: token})
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
