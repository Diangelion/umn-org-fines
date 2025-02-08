package utils

import (
	"backend/internal/models"
	"encoding/json"
	"net/http"
)

func SendJSONResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := models.APIResponse{
		Success: statusCode < 400,
		Message: message,
		Data: data,
	}

	json.NewEncoder(w).Encode(response)
}