package handlers

import (
	"encoding/json"
	"fmt"
	"gateway/internal/models"
	"gateway/internal/services"
	"log"
	"net/http"

	"gateway/utils"
)

var fileName string = "alert.html"

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	typeMsg := "registration"

	// Parse form data
	if err := utils.ParseRequestBody(r); err != nil {
		log.Println("RegisterUser | Parse request error: ", err)
		msg := fmt.Sprintf("Invalid input. %s", utils.LoginRegisterErrorMessage(&typeMsg))
		utils.SendAlert(w, "Error", msg, fileName)
		return
	}

	// Decode form data into User struct
	var user models.UserRegistration
	if err := utils.DecodeRequestBody(r, &user); err != nil {
		log.Println("RegisterUser | Decode request error: ", err)
		msg := fmt.Sprintf("Invalid %s form. %s", typeMsg, utils.LoginRegisterErrorMessage(&typeMsg))
		utils.SendAlert(w, "Error", msg, fileName)
		return
	}

	// Validate if required fields exist
	if user.Name == "" || user.Email == "" || user.Password == "" || user.ConfirmPassword == "" {
		log.Printf("RegisterUser | Missing required field(s)\n")
		msg := fmt.Sprintf("Missing required field(s). %s", utils.LoginRegisterErrorMessage(&typeMsg))
		utils.SendAlert(w, "Error", msg, fileName)
		return
	}

	// Validate password & confirm password
	if user.Password != user.ConfirmPassword {
		log.Printf("RegisterUser | Password & confirm password don't match\n")
		msg := fmt.Sprintf("Password & confirm password don't match. %s", utils.LoginRegisterErrorMessage(&typeMsg))
		utils.SendAlert(w, "Error", msg, fileName)
		return
	}

	// Create new User struct without confirm password
	forwardUser := models.ForwardUserRegistration{
		Name: user.Name,
		Email: user.Email,
		Password: user.Password,
	}

	// Forward registration request to the backend service
	response, err := services.ForwardUserRegistration(forwardUser)
	if err != nil { // This error means the request **did not reach** the backend (e.g., network failure)
		log.Println("RegisterUser | Forward registration error: ", err)
		utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName)
		return
	}
	defer response.Body.Close()

	// Read and decode JSON response
	var jsonResponse models.Response
	if err := json.NewDecoder(response.Body).Decode(&jsonResponse); err != nil {
		log.Println("RegisterUser | Decode response error: ", err)
		utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName)
		return
	}

	msg := jsonResponse.Message

	// Handle non-200 responses by overriding status code
	if response.StatusCode >= 400 {
		log.Print("RegisterUser | Response not ok: ", msg)
		if response.StatusCode != 409 { // If not conflict
			msg = utils.GetGeneralErrorMessage()
		}
		utils.SendAlert(w, "Failed", msg, fileName)
		return
	}

	w.Header().Set("HX-Trigger", "resetForm")
	utils.SendAlert(w, "Success", "Your account has been successfully created.", fileName)
}


func LoginUser(w http.ResponseWriter, r *http.Request) {
	typeMsg := "login"

	// Parse form data
	if err := utils.ParseRequestBody(r); err != nil {
		log.Println("LoginUser | Parse request error: ", err)
		msg := fmt.Sprintf("Invalid input. %s", utils.LoginRegisterErrorMessage(&typeMsg))
		utils.SendAlert(w, "Error", msg, fileName)
		return
	}

	// Decode form data into User struct
	var user models.UserLogin
	if err := utils.DecodeRequestBody(r, &user); err != nil {
		log.Println("LoginUser | Decode request error: ", err)
		msg :=  fmt.Sprintf("Invalid %s form. %s", typeMsg, utils.LoginRegisterErrorMessage(&typeMsg))
		utils.SendAlert(w, "Error", msg, fileName)
		return
	}

    // Validate if required fields exist
    if user.Email == "" || user.Password == "" {
		log.Printf("LoginUser | Missing required field(s)\n")
		msg := fmt.Sprintf("Missing required field(s). %s", utils.LoginRegisterErrorMessage(&typeMsg))
        utils.SendAlert(w, "Error", msg, fileName)
        return
    }

	// Forward registration request to the backend service
	response, err := services.ForwardUserLogin(user)
	if err != nil {
		log.Println("LoginUser | Forward login error: ", err)
        utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName)
		return
	}

	// Read and decode JSON response
	var jsonResponse models.Response
	if err := json.NewDecoder(response.Body).Decode(&jsonResponse); err != nil {
		log.Println("LoginUser | Decode response error: ", err)
		utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName)
		return
	}

	// Handle non-200 responses by overriding status code
	if response.StatusCode >= 400 {
		log.Print("LoginUser | Response not ok: ", jsonResponse.Message)
		utils.SendAlert(w, "Failed", jsonResponse.Message, fileName)
		return
	}

	// Generate an access token
	accessToken, err := utils.GenerateAccessToken(jsonResponse.Data["user_id"].(string))
	if err != nil {
		log.Print("LoginUser | Generate access token error: ", err)
		utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName)
		return
	}

	// Generate a refresh token
	refreshToken, err := utils.GenerateRefreshToken(jsonResponse.Data["user_id"].(string))
	if err != nil {
		log.Print("LoginUser | Generate refresh token error: ", err)
		utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName)
		return
	}

	// w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	// w.Header().Set("X-Refresh-Token", fmt.Sprintf("Bearer %s", refreshToken))
	authToken := models.AuthorizationToken{
		AccessToken: fmt.Sprintf("Bearer %s", accessToken),
		RefreshToken: fmt.Sprintf("Bearer %s", refreshToken),
	}

	jsonAuthToken, err := json.Marshal(authToken)
	if err != nil {
		log.Println("LoginUser | Marshal json auth token error: ", err)
		utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName)
	}

	w.Header().Set("HX-Trigger", fmt.Sprintf(`{"receiveJWT": %s}`, string(jsonAuthToken)))
	w.WriteHeader(http.StatusAccepted)
}

func EditUser(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(string)
	if !ok {
		log.Printf("EditUser | Context userId not found\n")
		utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName)
		return
	}

	typeMsg := "edit profile"

	// Parse form data
	if err := utils.ParseRequestBody(r); err != nil {
		log.Println("EditUser | Parse request error: ", err)
		msg := fmt.Sprintf("Invalid input. %s", utils.LoginRegisterErrorMessage(&typeMsg))
		utils.SendAlert(w, "Error", msg, fileName)
		return
	}

	// Decode form data into User struct
	var user models.UserEdit
	if err := utils.DecodeRequestBody(r, &user); err != nil {
		log.Println("EditUser | Decode request error: ", err)
		msg :=  fmt.Sprintf("Invalid %s form. %s", typeMsg, utils.LoginRegisterErrorMessage(&typeMsg))
		utils.SendAlert(w, "Error", msg, fileName)
		return
	}

	if user.Email == "" || user.Name == "" || user.ProfilePhoto == "" || user.CoverPhoto == "" {
		log.Printf("EditUser |Missing required field(s)\n")
		msg := fmt.Sprintf("Missing required field(s). %s", utils.LoginRegisterErrorMessage(&typeMsg))
		utils.SendAlert(w, "Error", msg, fileName)
		return
	}

	// Forward registration request to the backend service
	response, err := services.ForwardUserEdit(user, userId)
	if err != nil { // This error means the request **did not reach** the backend (e.g., network failure)
		log.Println("EditUser | Forward edit error: ", err)
		utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName)
		return
	}
	defer response.Body.Close()

	// Read and decode JSON response
	var jsonResponse models.Response
	if err := json.NewDecoder(response.Body).Decode(&jsonResponse); err != nil {
		log.Println("EditUser | Decode response error: ", err)
		utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName)
		return
	}

	// Handle non-200 responses by overriding status code
	if response.StatusCode >= 400 {
		log.Print("EditUser | Response not ok: ", jsonResponse.Message)
		utils.SendAlert(w, "Failed", jsonResponse.Message, fileName)
		return
	}

	utils.SendAlert(w, "Success", "Profile saved.", fileName)
}