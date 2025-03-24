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

type PagesHandler struct {
	Config *config.Config
}

func NewPagesHandler(cfg *config.Config) *PagesHandler {
    return &PagesHandler{Config: cfg}
}

func (h *PagesHandler) IndexPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("HX-Redirect", "/login")
	w.WriteHeader(http.StatusAccepted)
}

func (h *PagesHandler) RegisterPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("HX-Reswap", "innerHTML")
	w.Header().Set("HX-Retarget", "main")
	document := utils.GetAuthPage(h.Config.BaseURL)
	utils.SendAuthPage(w, document, "pages/register.html")
}

func (h *PagesHandler) LoginPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("HX-Reswap", "innerHTML")
	w.Header().Set("HX-Retarget", "main")
	document := utils.GetAuthPage(h.Config.BaseURL)
	utils.SendAuthPage(w, document, "pages/login.html")
}

func (h *PagesHandler) HomePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("HX-Reswap", "innerHTML")
	w.Header().Set("HX-Retarget", "main")
	document := utils.GetAuthPage(h.Config.BaseURL)
	utils.SendAuthPage(w, document, "pages/home.html")
}

func (h *PagesHandler) ProfilePage(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value(middleware.UserIdKey).(string)
	if !ok {
		log.Println("ProfilePage | Context userId not found")
		utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName)
		return
	}

	 // Define a simple result type for aggregating responses
    type result struct {
        Response *http.Response
        Error  error
    }

    // Channels to receive the responses
    userChan := make(chan result)
    orgChan := make(chan result)

    // Make concurrent calls to both backend services
    go func() {
        response, err := services.ForwardGetUser(userId)
        userChan <- result{Response: response, Error: err}
    }()
    go func() {
        response, err := services.ForwardGetListOrganization(userId)
        orgChan <- result{Response: response, Error: err}
    }()

    // Wait for both responses
    userRes := <-userChan
    orgRes := <-orgChan

    // Handle any error from either call
    if userRes.Error != nil || orgRes.Error != nil {
        log.Printf(`
			ProfilePage | Fetching user or organization data error\nUser error: %v\nOrganization error: %v
		`, userRes.Error, orgRes.Error)
        utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName)
        return
    }

	defer userRes.Response.Body.Close()
	defer orgRes.Response.Body.Close()

	// Read and decode user JSON response
	var userJsonResponse models.Response
	if err := json.NewDecoder(userRes.Response.Body).Decode(&userJsonResponse); err != nil {
		log.Println("ProfilePage | Decode user response error: ", err)
		utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName)
		return
	}

	// Handle non-200 responses by overriding status code
	if userRes.Response.StatusCode >= 400 {
		log.Print("ProfilePage | User response not ok: ", userJsonResponse.Message)
		utils.SendAlert(w, "Failed", userJsonResponse.Message, fileName)
		return
	}

	// Read and decode organization JSON response
	var orgJsonResponse models.Response
	if err := json.NewDecoder(orgRes.Response.Body).Decode(&orgJsonResponse); err != nil {
		log.Println("ProfilePage | Decode organization response error: ", err)
		utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName)
		return
	}

	// Handle non-200 responses by overriding status code
	if orgRes.Response.StatusCode >= 400 {
		log.Print("ProfilePage | Organization response not ok: ", userJsonResponse.Message)
		utils.SendAlert(w, "Failed", userJsonResponse.Message, fileName)
		return
	}

	baseURL :=  utils.GetAuthPage(h.Config.BaseURL).(models.AuthPage)

    // Combine the data
	user := &models.User{}
    if err := utils.CombineProfileAndOrganizations(user, userJsonResponse, orgJsonResponse); err != nil {
        log.Println("ProfilePage | Combining data error:", err)
        utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName)
        return
    }

	document := models.ProfilePage{
		AuthPage: baseURL,
		User: *user,
	}

	w.Header().Set("HX-Reswap", "innerHTML")
	w.Header().Set("HX-Retarget", "main")
	utils.SendAuthPage(w, document, "pages/profile.html")
}

func (h *PagesHandler) NotFoundPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("HX-Reswap", "innerHTML")
	w.Header().Set("HX-Retarget", "main")
	utils.SendHTMLDocumentResponse(w, nil, "pages/not_found.html")
}