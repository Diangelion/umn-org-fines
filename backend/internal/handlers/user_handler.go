package handlers

import (
	"backend/internal/models"
	"backend/internal/services"
	"backend/utils"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
    service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
    return &UserHandler{service}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
    var user models.UserRegistration

    // Assign request body to user for validation
    if errDecode := json.NewDecoder(r.Body).Decode(&user); errDecode != nil {
        utils.SendJSONResponse(w, http.StatusBadRequest, "invalid JSON payload", nil)
        return
    }

    if errRegister := h.service.RegisterUser(&user); errRegister != nil {
        utils.SendJSONResponse(w, http.StatusInternalServerError, errRegister.Error(), nil)
        return
    }

    utils.SendJSONResponse(w, http.StatusCreated, "Success: user created", user)
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
