package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gateway/config"
	"gateway/internal/models"
	"net/http"
)

func ForwardUserRegistration(user models.UserRegistration) (*http.Response, error) {
	cfg := config.LoadConfig()

	jsonData, err := json.Marshal(user)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal user data: %w", err)
	}

	registrationUrl := fmt.Sprintf("%s/user/register", cfg.BackendURL)
	contentType := "application/json"
	reqBody := bytes.NewBuffer(jsonData)

	client := &http.Client{}
	return client.Post(registrationUrl, contentType, reqBody)
}

func ForwardUserLogin(user models.UserLogin) (*http.Response, error) {
	cfg := config.LoadConfig()

	jsonData, err := json.Marshal(user)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal user data: %w", err)
	}

	loginUrl := fmt.Sprintf("%s/user/login", cfg.BackendURL)
	contentType := "application/json"
	reqBody := bytes.NewBuffer(jsonData)

	client := &http.Client{}
	return client.Post(loginUrl, contentType, reqBody)
}
