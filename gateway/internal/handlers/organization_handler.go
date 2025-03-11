package handlers

import (
	"fmt"
	"gateway/internal/models"
	"gateway/internal/services"
	"gateway/utils"
	"log"
	"net/http"
)

// func GetListOrganizations(w http.ResponseWriter, r *http.Request) {
// 	userId, ok := r.Context().Value("userId").(string)
// 	if !ok {
// 		log.Println("GetListOrganizations | Context userId not found")
// 		utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName)
// 		return
// 	}
// }

// func GetSingleOrganization(w http.ResponseWriter, r *http.Request) {
// 	userId, ok := r.Context().Value("userId").(string)
// 	if !ok {
// 		log.Println("GetSingleOrganization | Context userId not found")
// 		utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName)
// 		return
// 	}
// }

func CreateOrganizationHandler(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(string)
	if !ok {
		log.Println("CreateOrganizationHandler | Context userId not found")
		utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName)
		return
	}

	typeMsg := "create organization"

	// Parse form data
	if err := utils.ParseRequestBody(r); err != nil {
		log.Println("CreateOrganizationHandler | Parse request error: ", err)
		msg := fmt.Sprintf("Invalid input. %s", utils.InvalidFormErrorMessage(&typeMsg))
		utils.SendAlert(w, "Error", msg, fileName)
		return
	}

	// Decode form data into User struct
	var org models.CreateOrganization
	if err := utils.DecodeRequestBody(r, &org); err != nil {
		log.Println("CreateOrganizationHandler | Decode request error: ", err)
		msg := fmt.Sprintf("Invalid %s form. %s", typeMsg, utils.InvalidFormErrorMessage(&typeMsg))
		utils.SendAlert(w, "Error", msg, fileName)
		return
	}

	// Validate if required fields exist
	if org.OrganizationTitle == "" || org.OrganizationStartDate == "" || org.OrganizationEndDate == "" {
		log.Println("CreateOrganizationHandler | Missing required field(s)")
		msg := fmt.Sprintf("Missing required field(s). %s", utils.InvalidFormErrorMessage(&typeMsg))
		utils.SendAlert(w, "Error", msg, fileName)
		return
	}

	// Forward registration request to the backend service
	response, err := services.ForwardCreateOrganization(org, userId)
	if err != nil { // This error means the request **did not reach** the backend (e.g., network failure)
		log.Println("CreateOrganizationHandler | Forward registration error: ", err)
		utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName)
		return
	}
	defer response.Body.Close()

	// // Read and decode JSON response
	// var jsonResponse models.Response
	// if err := json.NewDecoder(response.Body).Decode(&jsonResponse); err != nil {
	// 	log.Println("CreateOrganizationHandler | Decode response error: ", err)
	// 	utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName)
	// 	return
	// }

	// msg := jsonResponse.Message

	// // Handle non-200 responses by overriding status code
	// if response.StatusCode >= 400 {
	// 	log.Print("CreateOrganizationHandler | Response not ok: ", msg)
	// 	if response.StatusCode != 409 { // If not conflict
	// 		msg = utils.GetGeneralErrorMessage()
	// 	}
	// 	utils.SendAlert(w, "Failed", msg, fileName)
	// 	return
	// }

	// w.Header().Set("HX-Trigger", "resetForm")
	// utils.SendAlert(w, "Success", "Your account has been successfully created.", fileName)
}

// func JoinOrganization(w http.ResponseWriter, r *http.Request) {
// }