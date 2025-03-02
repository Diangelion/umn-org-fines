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
    if err := utils.DecodeRequestBody(r, user); err != nil {
        log.Println("Register | Decode request error: ", err)
        utils.SendJSONResponse(w, http.StatusBadRequest, "Invalid JSON payload.", nil)
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

    utils.SendJSONResponse(w, http.StatusCreated, "Your account has been successfully created.", user)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
    var user models.UserLogin

    // Assign request body to user for validation
    if err := utils.DecodeRequestBody(r, user); err != nil {
        log.Println("Login | Decode request error: ", err)
        utils.SendJSONResponse(w, http.StatusBadRequest, "Invalid JSON payload.", nil)
        return
    }

    userId, err := h.service.LoginUser(&user);
    if err != nil {
        log.Println("Login | Login service error: ", err)
        utils.SendJSONResponse(w, http.StatusInternalServerError, err.Error(), nil)
        return
    }

    data := models.UserId{UserId: userId}
    utils.SendJSONResponse(w, http.StatusCreated, "Session created.", data)
}

func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
    userId := r.Header.Get("Authorization")

    if userId == "" {
        log.Println("Get | User id not found in Authorization")
        utils.SendJSONResponse(w, http.StatusBadRequest, "User id not found in request header.", nil)
        return
    }

    data, err := h.service.GetUser(userId)
    if err != nil {
    log.Println("Get | Get user error: ", err)
        utils.SendJSONResponse(w, http.StatusInternalServerError, err.Error(), nil)
        return
    }

    jsonData, err := json.Marshal(data)
    if err != nil {
        log.Println("Get | Marshaling user data to JSON error: ", err)
        utils.SendJSONResponse(w, http.StatusInternalServerError, "Error marshaling user data to JSON.", nil)
        return
    }

    utils.SendJSONResponse(w, http.StatusOK, "User found.", jsonData)
}


func (h *UserHandler) Edit(w http.ResponseWriter, r *http.Request) {
    userId := r.Header.Get("Authorization")
    if userId == "" {
        log.Println("Edit | User id not found in Authorization")
        utils.SendJSONResponse(w, http.StatusBadRequest, "User id not found in Authorization.", nil)
        return
    }

    var user models.UserEdit

    // Assign request body to user for validation
    if err := utils.DecodeRequestBody(r, user); err != nil {
        log.Println("Edit | Decode request error: ", err)
        utils.SendJSONResponse(w, http.StatusBadRequest, "Invalid JSON payload.", nil)
        return
    }

    if err := h.service.EditUser(&user, userId); err != nil {
        log.Println("Edit | Edit service error: ", err)
        utils.SendJSONResponse(w, http.StatusInternalServerError, err.Error(), nil)
        return
    }

    utils.SendJSONResponse(w, http.StatusOK, "Profile saved.", nil)
}