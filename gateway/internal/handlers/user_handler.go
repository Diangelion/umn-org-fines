package handlers

import (
	"encoding/json"
	"gateway/internal/models"
	"gateway/internal/services"
	"io"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := json.Unmarshal(body, &user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Forward registration request to the backend service
	response, err := services.ForwardUserRegistration(user)
	if err != nil {
		http.Error(w, "Error forwarding request to backend", http.StatusInternalServerError)
		return
	}

	// Return the response from the backend service
	w.WriteHeader(response.StatusCode)
	if response.Body != nil {
		defer response.Body.Close()
		body, _ := io.ReadAll(response.Body)
		w.Write(body)
	}
}


func LoginUser(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := json.Unmarshal(body, &user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Forward registration request to the backend service
	response, err := services.ForwardUserRegistration(user)
	if err != nil {
		http.Error(w, "Error forwarding request to backend", http.StatusInternalServerError)
		return
	}

	// Return the response from the backend service
	w.WriteHeader(response.StatusCode)
	if response.Body != nil {
		defer response.Body.Close()
		body, _ := io.ReadAll(response.Body)
		w.Write(body)
	}
}
