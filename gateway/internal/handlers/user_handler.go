package handlers

import (
	"encoding/json"
	"gateway/internal/services"
	"io"
	"io/ioutil"
	"net/http"
)

// User struct represents the user registration data
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// RegisterUser handles user registration
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	var user User
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
		body, _ := ioutil.ReadAll(response.Body)
		w.Write(body)
	}
}
