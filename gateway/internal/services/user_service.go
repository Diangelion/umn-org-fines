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

var cfg = config.LoadConfig()

func ForwardUserRegistration(user models.UserRegistration) (*http.Response, error) {
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Unable to marshal user data")
	}

	registrationUrl := fmt.Sprintf("%s/user/register", cfg.BackendURL)
	contentType := "application/json"
	reqBody := bytes.NewBuffer(jsonData)

	client := &http.Client{}
	return client.Post(registrationUrl, contentType, reqBody)
}

func ForwardUserLogin(user models.UserLogin) (*http.Response, error) {
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Unable to marshal user data")
	}

	loginUrl := fmt.Sprintf("%s/user/login", cfg.BackendURL)
	contentType := "application/json"
	reqBody := bytes.NewBuffer(jsonData)

	client := &http.Client{}
	return client.Post(loginUrl, contentType, reqBody)
}
