package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/achintha-dilshan/go-validator/utils/validator"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Name     string `json:"name" validate:"required,min=3"`
			Email    string `json:"email" validate:"required,email"`
			Password string `json:"password" validate:"required,min=8"`
		}

		// decode request body
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			fmt.Fprintf(w, "Invalid JSON payload.")
			return
		}

		// validate request
		validator := validator.New()
		if err := validator.Validate(req); err != nil {

			data, err := json.Marshal(err)
			if err != nil {
				log.Printf("Failed to marshal JSON response: %v", err)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(data)

			return
		}

		fmt.Fprintf(w, "Validation was successful!")
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Server is running on \"http://localhost:8080\"")

	if err := server.ListenAndServe(); err != nil {
		log.Println(err.Error())
	}
}
