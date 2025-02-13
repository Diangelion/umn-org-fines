package handlers

import (
	"encoding/json"
	"fmt"
	"gateway/internal/models"
	"gateway/internal/services"
	"io"
	"log"
	"net/http"

	"gateway/utils"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	typeMsg := "registration"
	fileName := "alert.html"

	// Parse form data
	if err := utils.ParseRequestBody(r); err != nil {
		msg := fmt.Sprintf("Invalid input. %s", utils.LoginRegisterErrorMessage(&typeMsg))
		utils.SendAlert(w, "Error", msg, fileName, http.StatusBadRequest)
		return
	}

	// Decode form data into User struct
	var user models.UserRegistration
	if err := utils.DecodeRequestBody(r, &user); err != nil {
		msg := fmt.Sprintf("Invalid %s form. %s", typeMsg, utils.LoginRegisterErrorMessage(&typeMsg))
		utils.SendAlert(w, "Error", msg, fileName, http.StatusBadRequest)
		return
	}

	// Validate if required fields exist
	if user.Name == "" || user.Email == "" || user.Password == "" {
		msg := fmt.Sprintf("Missing required field(s). %s", utils.LoginRegisterErrorMessage(&typeMsg))
		utils.SendAlert(w, "Error", msg, fileName, http.StatusBadRequest)
		return
	}
	
	// Forward registration request to the backend service
	response, err := services.ForwardUserRegistration(user)
	if err != nil { // This error means the request **did not reach** the backend (e.g., network failure)
		log.Println(err.Error())
		utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName, http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()
	
	// Read and decode JSON response
	var jsonResponse models.Response
	if err := json.NewDecoder(response.Body).Decode(&jsonResponse); err != nil {
		log.Println(err.Error())
		utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName, http.StatusInternalServerError)
		return
	}
	
	var alertMsg string
	var statusCode int = http.StatusOK // Make default success
	
	// Handle non-200 responses by overriding status code
	if response.StatusCode >= 400 {
		log.Print(jsonResponse.Message)
		alertMsg = "Failed"
		statusCode = http.StatusInternalServerError
	} else {
		alertMsg = "Success"
	}

	utils.SendAlert(w, alertMsg, jsonResponse.Message, fileName, statusCode)
}


func LoginUser(w http.ResponseWriter, r *http.Request) {
	typeMsg := "login"

	// Parse form data
	if err := utils.ParseRequestBody(r); err != nil {
		errParse := fmt.Sprintf("Error: invalid input. %s", utils.LoginRegisterErrorMessage(&typeMsg))
		http.Error(w, errParse, http.StatusBadRequest)
		return
	}

	// Decode form data into User struct
    var user models.UserLogin
	if err := utils.DecodeRequestBody(r, &user); err != nil {
		errDecode :=  fmt.Sprintf("Error: invalid %s form. %s", typeMsg, utils.LoginRegisterErrorMessage(&typeMsg))
		http.Error(w, errDecode, http.StatusBadRequest)
        return
	}
   
    // Validate if required fields exist
    if user.Email == "" || user.Password == "" {
		errMissing := fmt.Sprintf("Error: missing required field(s). %s", utils.LoginRegisterErrorMessage(&typeMsg))
        http.Error(w, errMissing, http.StatusBadRequest)
        return
    }

	// Forward registration request to the backend service
	response, err := services.ForwardUserLogin(user)
	if err != nil {
		http.Error(w, "Error: unable to process the request. Please try again later.", http.StatusInternalServerError)
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
