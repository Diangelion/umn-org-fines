package handlers

import (
	"fmt"
	"gateway/internal/models"
	"gateway/internal/services"
	"io"
	"net/http"

	"gateway/utils"
)

func generalErrorMessage(typeForm *string) string {
	return fmt.Sprintf("Please check again your %s form and try again.", *typeForm)
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	typeMsg := "registration"

	// Parse form data
	if err := utils.ParseRequestBody(r); err != nil {
		errParse := fmt.Sprintf("Invalid input. %s", generalErrorMessage(&typeMsg))
		documentData := utils.GetAlert("Error", errParse)
		utils.SendHTMLDocumentResponse(w, documentData, "alert.html", http.StatusBadRequest)
		return
	}

	// Decode form data into User struct
    var user models.UserRegistration
	if err := utils.DecodeRequestBody(r, &user); err != nil {
		errDecode :=  fmt.Sprintf("Invalid %s form. %s", typeMsg, generalErrorMessage(&typeMsg))
		documentData := utils.GetAlert("Error", errDecode)
		utils.SendHTMLDocumentResponse(w, documentData, "alert.html", http.StatusBadRequest)
        return
	}

    // Validate if required fields exist
    if user.Name == "" || user.Email == "" || user.Password == "" {
		errMissing := fmt.Sprintf("Missing required field(s). %s", generalErrorMessage(&typeMsg))
		documentData := utils.GetAlert("Error", errMissing)
		utils.SendHTMLDocumentResponse(w, documentData, "alert.html", http.StatusBadRequest)
        return
    }

	// Forward registration request to the backend service
	response, err := services.ForwardUserRegistration(user)
	if err != nil {
		documentData := utils.GetAlert("Error", "Unable to process the request. Please try again later.")
		utils.SendHTMLDocumentResponse(w, documentData, "alert.html", http.StatusInternalServerError)
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
	typeMsg := "login"

	// Parse form data
	if err := utils.ParseRequestBody(r); err != nil {
		errParse := fmt.Sprintf("Error: invalid input. %s", generalErrorMessage(&typeMsg))
		http.Error(w, errParse, http.StatusBadRequest)
		return
	}

	// Decode form data into User struct
    var user models.UserLogin
	if err := utils.DecodeRequestBody(r, &user); err != nil {
		errDecode :=  fmt.Sprintf("Error: invalid %s form. %s", typeMsg, generalErrorMessage(&typeMsg))
		http.Error(w, errDecode, http.StatusBadRequest)
        return
	}
   
    // Validate if required fields exist
    if user.Email == "" || user.Password == "" {
		errMissing := fmt.Sprintf("Error: missing required field(s). %s", generalErrorMessage(&typeMsg))
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
