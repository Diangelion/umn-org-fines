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
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        log.Println("Register | Decode request error: ", err)
        utils.SendJSONResponse(w, http.StatusBadRequest, "Invalid JSON payload", nil)
        return
    }

    // Perform registration
    if err := h.service.RegisterUser(user); err != nil {
        log.Println("Register | Registration service error: ", err)

        // Handle duplicate email specifically, differentiate with response status code
        var dupErr *models.DuplicateEmailError
        statusCode := utils.StatusCodeForError(err, dupErr, http.StatusConflict)
        utils.SendJSONResponse(w, statusCode, err.Error(), nil)
        return
    }

    utils.SendJSONResponse(w, http.StatusCreated, "User created", user)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
    var user models.UserLogin
    
    // Assign request body to user for validation
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        log.Println(err)
        utils.SendJSONResponse(w, http.StatusBadRequest, "Invalid JSON payload", nil)
        return
    }

    userId, err := h.service.LoginUser(&user);
    if err != nil {
        utils.SendJSONResponse(w, http.StatusInternalServerError, err.Error(), nil)
        return
    }

    data := models.UserId{UserId: userId}
    utils.SendJSONResponse(w, http.StatusCreated, "Session created", data)
}
