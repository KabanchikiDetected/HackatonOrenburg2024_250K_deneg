package main

import (
	"crypto/rsa"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/KabanchikiDetected/users/pkg/users"
	"github.com/golang-jwt/jwt/v5"
)

func getKey() *rsa.PublicKey {
	data, err := os.ReadFile("./keys/example_public_key.pem")
	if err != nil {
		log.Fatal(err)
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(data)
	if err != nil {
		log.Fatal(err)
	}

	return key
}

func main() {
	// Get public Key
	key := getKey()

	// Create middleware with public key
	mw := users.MiddlwareJWT(key)

	// Index page.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./testServer/index.html")
	})

	// Unauthorized if invalid token and hello message if all ok.
	http.Handle("/me", mw(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload, err := users.FromContext(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Cool marshalling. :xd:
		fmt.Fprintf(w, "{\"msg\": \"Hello %s\"}", payload.ID)
		w.WriteHeader(http.StatusOK)
	})))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
