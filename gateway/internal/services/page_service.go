package services

import (
	"errors"
	"fmt"
	"gateway/config"
	"log"
	"net/http"
)

func ForwardProfilePage(userId string) (*http.Response, error) {
	cfg := config.LoadConfig()

	registrationUrl := fmt.Sprintf("%s/user/get", cfg.BackendURL)
	contentType := "application/json"

	// Create a new HTTP request with the JSON data
	req, err := http.NewRequest("GET", registrationUrl, nil)
	if err != nil {
		log.Println("ForwardProfilePage | Create request error: ", err)
		return nil, errors.New("Unable to create HTTP request.")
	}

	// Set the Content-Type header
	req.Header.Set("Content-Type", contentType)
	// Set the Authorization header (e.g., using a Bearer token)
	req.Header.Set("Authorization", userId)

	// Send the request using http.Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("ForwardProfilePage | Send request error: ", err)
		return nil, errors.New("Unable to send HTTP request.")
	}

	// Return the response
	return resp, nil
}