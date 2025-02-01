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
		return nil, fmt.Errorf("Failed to marshal user data: %w", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/user/register", cfg.BackendURL), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	return client.Do(req)
}
