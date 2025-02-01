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
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Failed: invalid JSON payload", http.StatusBadRequest)
        return
    }

    if err := h.service.RegisterUser(&user); err != nil {
        http.Error(w, "Failed: cannot register user.", http.StatusInternalServerError)
        return
    }

    utils.SendJSONResponse(w, http.StatusCreated, "Success: user created", user)
}
