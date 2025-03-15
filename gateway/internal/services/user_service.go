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

func ForwardRegisterUser(user models.ForwardRegisterUser) (*http.Response, error) {
	cfg := config.LoadConfig()

	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Println("ForwardRegisterUser | Marshal error: ", err)
		return nil, errors.New("Unable to marshal user data")
	}

	registrationUrl := fmt.Sprintf("%s/user/register", cfg.BackendURL)
	contentType := "application/json"
	reqBody := bytes.NewBuffer(jsonData)

	client := &http.Client{}
	return client.Post(registrationUrl, contentType, reqBody)
}

func ForwardLoginUser(user models.LoginUser) (*http.Response, error) {
	cfg := config.LoadConfig()

	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Println("ForwardLoginUser | Marshal error: ", err)
		return nil, errors.New("Unable to marshal user data")
	}

	loginUrl := fmt.Sprintf("%s/user/login", cfg.BackendURL)
	contentType := "application/json"
	reqBody := bytes.NewBuffer(jsonData)

	client := &http.Client{}
	return client.Post(loginUrl, contentType, reqBody)
}

func ForwardGetUser(userId string) (*http.Response, error) {
	cfg := config.LoadConfig()

	registrationUrl := fmt.Sprintf("%s/user/get", cfg.BackendURL)

	// Create a new HTTP request with the JSON data
	req, err := http.NewRequest("GET", registrationUrl, nil)
	if err != nil {
		log.Println("ForwardGetUser | Create request error: ", err)
		return nil, errors.New("Unable to create HTTP request.")
	}

	// Set the Authorization header (e.g., using a Bearer token)
	req.Header.Set("Authorization", userId)

	// Send the request using http.Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("ForwardGetUser | Send request error: ", err)
		return nil, errors.New("Unable to send HTTP request.")
	}

	// Return the response
	return resp, nil
}

func ForwardEditUser(user models.EditUser, userId string) (*http.Response, error) {
	cfg := config.LoadConfig()

	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Println("ForwardEditUser | Marshal error: ", err)
		return nil, errors.New("Unable to marshal user data")
	}

	registrationUrl := fmt.Sprintf("%s/user/edit", cfg.BackendURL)
	contentType := "application/json"
	reqBody := bytes.NewBuffer(jsonData)

	// Create a new HTTP request with the JSON data
	req, err := http.NewRequest("POST", registrationUrl, reqBody)
	if err != nil {
		log.Println("ForwardEditUser | Create request error: ", err)
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
		log.Println("ForwardEditUser | Send request error: ", err)
		return nil, errors.New("Unable to send HTTP request.")
	}

	// Return the response
	return resp, nil
}