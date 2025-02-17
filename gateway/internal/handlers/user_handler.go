package handlers

import (
	"encoding/json"
	"fmt"
	"gateway/internal/models"
	"gateway/internal/services"
	"log"
	"net/http"
	"time"

	"gateway/utils"
)

var fileName string = "alert.html"

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	typeMsg := "registration"

	// Parse form data
	if err := utils.ParseRequestBody(r); err != nil {
		log.Println("RegisterUser | Parse request error: ", err)
		msg := fmt.Sprintf("Invalid input. %s", utils.LoginRegisterErrorMessage(&typeMsg))
		utils.SendAlert(w, "Error", msg, fileName, http.StatusInternalServerError)
		return
	}

	// Decode form data into User struct
	var user models.UserRegistration
	if err := utils.DecodeRequestBody(r, &user); err != nil {
		log.Println("RegisterUser | Decode request error: ", err)
		msg := fmt.Sprintf("Invalid %s form. %s", typeMsg, utils.LoginRegisterErrorMessage(&typeMsg))
		utils.SendAlert(w, "Error", msg, fileName, http.StatusInternalServerError)
		return
	}

	// Validate if required fields exist
	if user.Name == "" || user.Email == "" || user.Password == "" {
		log.Printf("Missing required field(s)")
		msg := fmt.Sprintf("Missing required field(s). %s", utils.LoginRegisterErrorMessage(&typeMsg))
		utils.SendAlert(w, "Error", msg, fileName, http.StatusInternalServerError)
		return
	}
	
	// Forward registration request to the backend service
	response, err := services.ForwardUserRegistration(user)
	if err != nil { // This error means the request **did not reach** the backend (e.g., network failure)
		log.Println("RegisterUser | Forward registration error: ", err)
		utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName, http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()
	
	// Read and decode JSON response
	var jsonResponse models.Response
	if err := json.NewDecoder(response.Body).Decode(&jsonResponse); err != nil {
		log.Println("RegisterUser | Decode response error: ", err)
		utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName, http.StatusInternalServerError)
		return
	}
	
	// Handle non-200 responses by overriding status code
	if response.StatusCode >= 400 {
		log.Print("RegisterUser | Response not ok: ", jsonResponse.Message)
		utils.SendAlert(w, "Failed", utils.GetGeneralErrorMessage(), fileName, http.StatusInternalServerError)
		return
		} 
		
		utils.SendAlert(w, "Success", jsonResponse.Message, fileName, http.StatusOK)
	}
	
	
	func LoginUser(w http.ResponseWriter, r *http.Request) {
		typeMsg := "login"
		
		// Parse form data
		if err := utils.ParseRequestBody(r); err != nil {
			log.Println("LoginUser | Parse request error: ", err)
			msg := fmt.Sprintf("Invalid input. %s", utils.LoginRegisterErrorMessage(&typeMsg))
			utils.SendAlert(w, "Error", msg, fileName, http.StatusInternalServerError)
			return
		}
		
		// Decode form data into User struct
		var user models.UserLogin
		if err := utils.DecodeRequestBody(r, &user); err != nil {
			log.Println("LoginUser | Decode request error: ", err)
			msg :=  fmt.Sprintf("Invalid %s form. %s", typeMsg, utils.LoginRegisterErrorMessage(&typeMsg))
			utils.SendAlert(w, "Error", msg, fileName, http.StatusInternalServerError)
			return
		}
		
    // Validate if required fields exist
    if user.Email == "" || user.Password == "" {
		log.Printf("Missing required field(s)")
		msg := fmt.Sprintf("Missing required field(s). %s", utils.LoginRegisterErrorMessage(&typeMsg))
        utils.SendAlert(w, "Error", msg, fileName, http.StatusInternalServerError)
        return
    }
	
	// Forward registration request to the backend service
	response, err := services.ForwardUserLogin(user)
	if err != nil {
		log.Println("LoginUser | Forward login error: ", err)
        utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName, http.StatusInternalServerError)
		return
	}
	
	// Read and decode JSON response
	var jsonResponse models.Response
	if err := json.NewDecoder(response.Body).Decode(&jsonResponse); err != nil {
		log.Println("LoginUser | Decode response error: ", err)
		utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName, http.StatusInternalServerError)
		return
	}
	
	// Handle non-200 responses by overriding status code
	if response.StatusCode >= 400 {
		log.Print("LoginUser | Response not ok: ", jsonResponse.Message)
		utils.SendAlert(w, "Failed", jsonResponse.Message, fileName, http.StatusInternalServerError)
		return
	} 
	
	// Generate an access token
	accessToken, err := utils.GenerateAccessToken(jsonResponse.Data["user_id"].(string))
	if err != nil {
		log.Print("LoginUser | Generate access token error: ", err)
		utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName, http.StatusInternalServerError)
		return
	}
	
	// Generate a refresh token
	refreshToken, err := utils.GenerateRefreshToken(jsonResponse.Data["user_id"].(string))
	if err != nil {
		log.Print("LoginUser | Generate refresh token error: ", err)
		utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName, http.StatusInternalServerError)
		return
	}

	// Set the tokens as HTTP-only cookies
    http.SetCookie(w, &http.Cookie{
        Name:     "access_token",
        Value:    accessToken,
        HttpOnly: true,
        Secure:   false, // set to true if using HTTPS
        Path:     "/",
        Expires:  time.Now().Add(15 * time.Minute), // same as access token expiry
    })
    http.SetCookie(w, &http.Cookie{
        Name:     "refresh_token",
        Value:    refreshToken,
        HttpOnly: true,
        Secure:   false, // set to true if using HTTPS
        Path:     "/",
        Expires:  time.Now().Add(1 * 24 * time.Hour), // same as refresh token expiry
    })
	
	w.Header().Set("HX-Target", "main")
	w.Header().Set("HX-Swap", "innerHTML")
	w.Header().Add("Access-Control-Expose-Headers", "HX")
	w.WriteHeader(http.StatusOK)
}


func IsLoggedIn(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "<div></div>")
}