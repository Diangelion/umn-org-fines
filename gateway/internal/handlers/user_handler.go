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

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	typeMsg := "registration"

	// Parse form data
	if err := utils.ParseRequestBody(r); err != nil {
		log.Println("RegisterUserHandler | Parse request error: ", err)
		msg := fmt.Sprintf("Invalid input. %s", utils.InvalidFormErrorMessage(&typeMsg))
		utils.SendAlert(w, "Error", msg, fileName)
		return
	}

	// Decode form data into User struct
	var user models.RegisterUser
	if err := utils.DecodeRequestBody(r, &user); err != nil {
		log.Println("RegisterUserHandler | Decode request error: ", err)
		msg := fmt.Sprintf("Invalid %s form. %s", typeMsg, utils.InvalidFormErrorMessage(&typeMsg))
		utils.SendAlert(w, "Error", msg, fileName)
		return
	}

	// Validate if required fields exist
	if user.Name == "" || user.Email == "" || user.Password == "" || user.ConfirmPassword == "" {
		log.Println("RegisterUserHandler | Missing required field(s)")
		msg := fmt.Sprintf("Missing required field(s). %s", utils.InvalidFormErrorMessage(&typeMsg))
		utils.SendAlert(w, "Error", msg, fileName)
		return
	}

	// Validate password & confirm password
	if user.Password != user.ConfirmPassword {
		log.Println("RegisterUserHandler | Password & confirm password don't match")
		msg := fmt.Sprintf("Password & confirm password don't match. %s", utils.InvalidFormErrorMessage(&typeMsg))
		utils.SendAlert(w, "Error", msg, fileName)
		return
	}

	// Create new User struct without confirm password
	forwardUser := models.ForwardRegisterUser{
		Name: user.Name,
		Email: user.Email,
		Password: user.Password,
	}

	// Forward registration request to the backend service
	response, err := services.ForwardRegisterUser(forwardUser)
	if err != nil { // This error means the request **did not reach** the backend (e.g., network failure)
		log.Println("RegisterUserHandler | Forward registration error: ", err)
		utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName)
		return
	}
	defer response.Body.Close()

	// Read and decode JSON response
	var jsonResponse models.Response
	if err := json.NewDecoder(response.Body).Decode(&jsonResponse); err != nil {
		log.Println("RegisterUserHandler | Decode response error: ", err)
		utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName)
		return
	}

	msg := jsonResponse.Message

	// Handle non-200 responses by overriding status code
	if response.StatusCode >= 400 {
		log.Print("RegisterUserHandler | Response not ok: ", msg)
		if response.StatusCode != 409 { // If not conflict
			msg = utils.GetGeneralErrorMessage()
		}
		utils.SendAlert(w, "Failed", msg, fileName)
		return
	}

	w.Header().Set("HX-Trigger", "resetForm")
	utils.SendAlert(w, "Success", "Your account has been successfully created.", fileName)
}


func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	typeMsg := "login"

	// Parse form data
	if err := utils.ParseRequestBody(r); err != nil {
		log.Println("LoginUserHandler | Parse request error: ", err)
		msg := fmt.Sprintf("Invalid input. %s", utils.InvalidFormErrorMessage(&typeMsg))
		utils.SendAlert(w, "Error", msg, fileName)
		return
	}

	// Decode form data into User struct
	var user models.LoginUser
	if err := utils.DecodeRequestBody(r, &user); err != nil {
		log.Println("LoginUserHandler | Decode request error: ", err)
		msg :=  fmt.Sprintf("Invalid %s form. %s", typeMsg, utils.InvalidFormErrorMessage(&typeMsg))
		utils.SendAlert(w, "Error", msg, fileName)
		return
	}

    // Validate if required fields exist
    if user.Email == "" || user.Password == "" {
		log.Println("LoginUserHandler | Missing required field(s)")
		msg := fmt.Sprintf("Missing required field(s). %s", utils.InvalidFormErrorMessage(&typeMsg))
        utils.SendAlert(w, "Error", msg, fileName)
        return
    }

	// Forward registration request to the backend service
	response, err := services.ForwardLoginUser(user)
	if err != nil {
		log.Println("LoginUserHandler | Forward login error: ", err)
        utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName)
		return
	}

	// Read and decode JSON response
	var jsonResponse models.Response
	if err := json.NewDecoder(response.Body).Decode(&jsonResponse); err != nil {
		log.Println("LoginUserHandler | Decode response error: ", err)
		utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName)
		return
	}

	// Handle non-200 responses by overriding status code
	if response.StatusCode >= 400 {
		log.Print("LoginUserHandler | Response not ok: ", jsonResponse.Message)
		utils.SendAlert(w, "Failed", jsonResponse.Message, fileName)
		return
	}

	// Generate an access token
	accessToken, err := utils.GenerateAccessToken(jsonResponse.Data["user_id"].(string))
	if err != nil {
		log.Print("LoginUserHandler | Generate access token error: ", err)
		utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName)
		return
	}

	// Generate a refresh token
	refreshToken, err := utils.GenerateRefreshToken(jsonResponse.Data["user_id"].(string))
	if err != nil {
		log.Print("LoginUserHandler | Generate refresh token error: ", err)
		utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName)
		return
	}

	authToken := models.AuthorizationToken{
		AccessToken: fmt.Sprintf("Bearer %s", accessToken),
		RefreshToken: fmt.Sprintf("Bearer %s", refreshToken),
	}

	jsonAuthToken, err := json.Marshal(authToken)
	if err != nil {
		log.Println("LoginUserHandler | Marshal json auth token error: ", err)
		utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName)
	}

	w.Header().Set("HX-Trigger", fmt.Sprintf(`{"receiveJWT": %s}`, string(jsonAuthToken)))
	w.WriteHeader(http.StatusAccepted)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	// userId, ok := r.Context().Value("userId").(string)
	// if !ok {
	// 	log.Println("GetUserHandler | Context userId not found")
	// 	utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName)
	// 	return
	// }

	// response, err := services.ForwardUserGet(userId)
	// if err != nil { // This error means the request **did not reach** the backend (e.g., network failure)
	// 	log.Println("GetUserHandler | Forward user get error: ", err)
	// 	utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName)
	// 	return
	// }
	// defer response.Body.Close()

	// // Read and decode JSON response
	// var jsonResponse models.Response
	// if err := json.NewDecoder(response.Body).Decode(&jsonResponse); err != nil {
	// 	log.Println("EditUser | Decode response error: ", err)
	// 	utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName)
	// 	return
	// }

	// // Handle non-200 responses by overriding status code
	// if response.StatusCode >= 400 {
	// 	log.Print("EditUser | Response not ok: ", jsonResponse.Message)
	// 	utils.SendAlert(w, "Failed", jsonResponse.Message, fileName)
	// 	return
	// }

	// utils.SendAlert(w, "Success", "Profile saved.", fileName)
}

func EditUserHandler(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(string)
	if !ok {
		log.Println("EditUserHandler | Context userId not found")
		utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName)
		return
	}

	typeMsg := "edit profile"

	// Parse form data
	if err := utils.ParseRequestBody(r); err != nil {
		log.Println("EditUserHandler | Parse request error: ", err)
		msg := fmt.Sprintf("Invalid input. %s", utils.InvalidFormErrorMessage(&typeMsg))
		utils.SendAlert(w, "Error", msg, fileName)
		return
	}

	// Decode form data into User struct
	var user models.EditUser
	if err := utils.DecodeRequestBody(r, &user); err != nil {
		log.Println("EditUserHandler | Decode request error: ", err)
		msg :=  fmt.Sprintf("Invalid %s form. %s", typeMsg, utils.InvalidFormErrorMessage(&typeMsg))
		utils.SendAlert(w, "Error", msg, fileName)
		return
	}

	if user.Email == "" || user.Name == "" || user.ProfilePhoto == "" || user.CoverPhoto == "" {
		log.Println("EditUserHandler | Missing required field(s)")
		msg := fmt.Sprintf("Missing required field(s). %s", utils.InvalidFormErrorMessage(&typeMsg))
		utils.SendAlert(w, "Error", msg, fileName)
		return
	}

	// Forward registration request to the backend service
	response, err := services.ForwardEditUser(user, userId)
	if err != nil { // This error means the request **did not reach** the backend (e.g., network failure)
		log.Println("EditUserHandler | Forward edit error: ", err)
		utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName)
		return
	}
	defer response.Body.Close()

	// Read and decode JSON response
	var jsonResponse models.Response
	if err := json.NewDecoder(response.Body).Decode(&jsonResponse); err != nil {
		log.Println("EditUserHandler | Decode response error: ", err)
		utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName)
		return
	}

	// Handle non-200 responses by overriding status code
	if response.StatusCode >= 400 {
		log.Print("EditUserHandler | Response not ok: ", jsonResponse.Message)
		utils.SendAlert(w, "Failed", jsonResponse.Message, fileName)
		return
	}

	utils.SendAlert(w, "Success", "Profile saved.", fileName)
}