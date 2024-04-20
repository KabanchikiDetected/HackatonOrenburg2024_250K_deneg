package utils

import (
	"encoding/json"
	"errors"
	"net/http"

	customErrors "github.com/KabanchikiDetected/hackaton/events/internal/errors"
	"strings"
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
}
