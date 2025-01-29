package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"umn-org-fines/internal/models"
	"umn-org-fines/internal/services"
)

type UserHandler struct {
    service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
    return &UserHandler{service}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Jalan")
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    if err := h.service.RegisterUser(&user); err != nil {
        http.Error(w, "User registration failed", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}
