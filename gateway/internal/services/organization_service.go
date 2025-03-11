package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gateway/config"
	"gateway/internal/models"
	"log"
	"net/http"
)

func ForwardCreateOrganization(org models.CreateOrganization, userId string) (*http.Response, error) {
	cfg := config.LoadConfig()

	jsonData, err := json.Marshal(org)
	if err != nil {
		log.Println("ForwardCreateOrganization | Marshal error: ", err)
		return nil, errors.New("Unable to marshal organization data")
	}

	registrationUrl := fmt.Sprintf("%s/org/create", cfg.BackendURL)
	contentType := "application/json"
	reqBody := bytes.NewBuffer(jsonData)

	// Create a new HTTP request with the JSON data
	req, err := http.NewRequest("POST", registrationUrl, reqBody)
	if err != nil {
		log.Println("ForwardCreateOrganization | Create request error: ", err)
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
		log.Println("ForwardCreateOrganization | Send request error: ", err)
		return nil, errors.New("Unable to send HTTP request.")
	}

	// Return the response
	return resp, nil
}