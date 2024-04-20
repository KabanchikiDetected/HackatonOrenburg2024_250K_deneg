package utils

import (
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"strings"

	customErrors "github.com/KabanchikiDetected/hackaton/students/internal/errors"
	"github.com/golang-jwt/jwt"
)

func Encode(w http.ResponseWriter, r *http.Request, v interface{}) error {
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	return err
}

func Decode(w http.ResponseWriter, r *http.Request, v interface{}) error {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
	return err
}

func SendResponceMessage(w http.ResponseWriter, message string) {
	resp := make(map[string]string)
	resp["message"] = message

	json.NewEncoder(w).Encode(resp)
}

func GetIdFromPath(w http.ResponseWriter, r *http.Request) string {
	pathID := r.PathValue("id")
	return pathID
}

func GetBoolQuery(value string) bool {
	return strings.ToLower(value) == "true"
}

func SendErrorMessage(w http.ResponseWriter, message string) {
	resp := make(map[string]string)
	resp["error"] = message

	json.NewEncoder(w).Encode(resp)
}

func HandleError(err error, w http.ResponseWriter) {
	if errors.Is(err, customErrors.InternalServerError) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if errors.Is(err, customErrors.NotFound) {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	if errors.Is(err, customErrors.BadRequest) {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if errors.Is(err, customErrors.Forbidden) {
		http.Error(w, err.Error(), http.StatusForbidden)
	}
}

func GetKey() *rsa.PublicKey {
	data, err := os.ReadFile("./keys/public_key.pem")
	if err != nil {
		fmt.Printf("Error reading public key: %v", err)
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(data)
	if err != nil {
		fmt.Printf("Error parsing public key: %v", err)
	}

	return key
}
