package handlers

import (
	"backend/internal/models"
	"backend/internal/services"
	"backend/utils"
	"encoding/json"
	"log"
	"net/http"
)

type UserHandler struct {
    service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
    return &UserHandler{service}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
    var user *models.UserRegistration

    // Assign request body to user for validation
    if errDecode := json.NewDecoder(r.Body).Decode(&user); errDecode != nil {
        log.Println(errDecode.Error())
        utils.SendJSONResponse(w, http.StatusBadRequest, "Invalid JSON payload", nil)
        return
    }

    // Perform registration
    if errRegister := h.service.RegisterUser(user); errRegister != nil {
        // Handle duplicate email specifically, differentiate with response status code
        var dupErr *models.DuplicateEmailError
        statusCode := utils.StatusCodeForError(errRegister, dupErr, http.StatusConflict)
        utils.SendJSONResponse(w, statusCode, errRegister.Error(), nil)
        return
    }

    utils.SendJSONResponse(w, http.StatusCreated, "User created", user)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
    var user models.UserLogin
    
    // Assign request body to user for validation
    if errDecode := json.NewDecoder(r.Body).Decode(&user); errDecode != nil {
        utils.SendJSONResponse(w, http.StatusBadRequest, "invalid JSON payload", nil)
        return
    }

    if errLogin := h.service.LoginUser(&user); errLogin != nil {
        utils.SendJSONResponse(w, http.StatusInternalServerError, errLogin.Error(), nil)
        return
    }

    utils.SendJSONResponse(w, http.StatusCreated, "session created", user)
}
