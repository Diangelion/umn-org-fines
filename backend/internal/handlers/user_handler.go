package handlers

import (
	"backend/internal/models"
	"backend/internal/services"
	"backend/utils"
	"log"
	"net/http"
)

type UserHandler struct {
    service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
    return &UserHandler{service}
}

func (h *UserHandler) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
    var user *models.RegisterUser

    // Assign request body to user for validation
    if err := utils.DecodeRequestBody(r, user); err != nil {
        log.Println("RegisterUserHandler | Decode request error: ", err)
        utils.SendJSONResponse(w, http.StatusBadRequest, "Invalid JSON payload.", nil)
        return
    }

    // Perform registration
    if err := h.service.RegisterUserService(user); err != nil {
        log.Println("RegisterUserHandler | Registration service error: ", err)

        // Handle duplicate email specifically, differentiate with response status code
        var dupErr *models.DuplicateEmailError
        statusCode := utils.StatusCodeForError(err, dupErr, http.StatusConflict)
        utils.SendJSONResponse(w, statusCode, err.Error(), nil)
        return
    }

    utils.SendJSONResponse(w, http.StatusCreated, "Your account has been successfully created.", user)
}

func (h *UserHandler) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
    var user models.LoginUser

    // Assign request body to user for validation
    if err := utils.DecodeRequestBody(r, user); err != nil {
        log.Println("LoginUserHandler | Decode request error: ", err)
        utils.SendJSONResponse(w, http.StatusBadRequest, "Invalid JSON payload.", nil)
        return
    }

    userId, err := h.service.LoginUserService(&user);
    if err != nil {
        log.Println("LoginUserHandler | Login service error: ", err)
        utils.SendJSONResponse(w, http.StatusInternalServerError, err.Error(), nil)
        return
    }

    data := models.UserId{UserId: userId}
    utils.SendJSONResponse(w, http.StatusCreated, "Session created.", data)
}

func (h *UserHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
    // userId := r.Header.Get("Authorization")

    // if userId == "" {
    //     log.Println("Get | User id not found in Authorization")
    //     utils.SendJSONResponse(w, http.StatusBadRequest, "User id not found in request header.", nil)
    //     return
    // }

    // data, err := h.service.GetUserService(userId)
    // if err != nil {
    // log.Println("Get | Get user error: ", err)
    //     utils.SendJSONResponse(w, http.StatusInternalServerError, err.Error(), nil)
    //     return
    // }

    // jsonData, err := json.Marshal(data)
    // if err != nil {
    //     log.Println("Get | Marshaling user data to JSON error: ", err)
    //     utils.SendJSONResponse(w, http.StatusInternalServerError, "Error marshaling user data to JSON.", nil)
    //     return
    // }

    // utils.SendJSONResponse(w, http.StatusOK, "User found.", jsonData)
}


func (h *UserHandler) EditUserHandler(w http.ResponseWriter, r *http.Request) {
    userId := r.Header.Get("Authorization")
    if userId == "" {
        log.Println("EditUserHandler | User id not found in Authorization")
        utils.SendJSONResponse(w, http.StatusBadRequest, "User id not found in Authorization.", nil)
        return
    }

    var user models.EditUser

    // Assign request body to user for validation
    if err := utils.DecodeRequestBody(r, user); err != nil {
        log.Println("EditUserHandler | Decode request error: ", err)
        utils.SendJSONResponse(w, http.StatusBadRequest, "Invalid JSON payload.", nil)
        return
    }

    if err := h.service.EditUserService(&user, userId); err != nil {
        log.Println("EditUserHandler | Edit service error: ", err)
        utils.SendJSONResponse(w, http.StatusInternalServerError, err.Error(), nil)
        return
    }

    utils.SendJSONResponse(w, http.StatusOK, "Profile saved.", nil)
}