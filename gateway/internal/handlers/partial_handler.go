package handlers

import (
	"encoding/json"
	"gateway/config"
	"gateway/internal/models"
	"gateway/internal/services"
	"gateway/middleware"
	"gateway/utils"
	"log"
	"net/http"
)

type PartialHandler struct {
	Config *config.Config
}

func NewPartialHandler(cfg *config.Config) *PartialHandler {
    return &PartialHandler{Config: cfg}
}

func (h *PartialHandler) SidebarProfilePartial(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value(middleware.UserIdKey).(string)
	if !ok {
		log.Println("SidebarProfilePartial | Context userId not found")
		utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName)
		return
	}

	response, err := services.ForwardGetUser(userId)
	if err != nil { // This error means the request **did not reach** the backend (e.g., network failure)
		log.Println("SidebarProfilePartial | Forward user get error: ", err)
		utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName)
		return
	}
	defer response.Body.Close()

	// Read and decode JSON response
	var jsonResponse models.Response
	if err := json.NewDecoder(response.Body).Decode(&jsonResponse); err != nil {
		log.Println("SidebarProfilePartial | Decode response error: ", err)
		utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName)
		return
	}

	// Handle non-200 responses by overriding status code
	if response.StatusCode >= 400 {
		log.Print("SidebarProfilePartial | Response not ok: ", jsonResponse.Message)
		utils.SendAlert(w, "Failed", jsonResponse.Message, fileName)
		return
	}

	document := models.EditUser{
		Name: jsonResponse.Data["name"].(string),
		Email: jsonResponse.Data["email"].(string),
		ProfilePhoto: jsonResponse.Data["profile_photo"].(string),
		CoverPhoto: jsonResponse.Data["cover_photo"].(string),
	}
	utils.SendAuthPartial(w, document, "partials/sidebar_profile.html")
}

func (h *PartialHandler) SidebarOrganizationListPartial(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value(middleware.UserIdKey).(string)
	if !ok {
		log.Println("SidebarOrganizationListPartial | Context userId not found")
		utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName)
		return
	}

	response, err := services.ForwardGetListOrganization(userId)
	if err != nil { // This error means the request **did not reach** the backend (e.g., network failure)
		log.Println("SidebarOrganizationListPartial | Forward user get error: ", err)
		utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName)
		return
	}
	defer response.Body.Close()

	// Read and decode JSON response
	var jsonResponse models.Response
	if err := json.NewDecoder(response.Body).Decode(&jsonResponse); err != nil {
		log.Println("SidebarOrganizationListPartial | Decode response error: ", err)
		utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName)
		return
	}

	// Handle non-200 responses by overriding status code
	if response.StatusCode >= 400 {
		log.Print("SidebarOrganizationListPartial | Response not ok: ", jsonResponse.Message)
		utils.SendAlert(w, "Failed", jsonResponse.Message, fileName)
		return
	}

	listData, ok := jsonResponse.Data["list"].([]string)
	if !ok {
		listData = []string{} // Ensure it's an empty slice, not nil
	}

	document := models.OrganizationList{
		List: listData,
	}
	utils.SendAuthPartial(w, document, "partials/sidebar_organization_list.html")
}