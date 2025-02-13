package utils

import (
	"backend/internal/models"
	"encoding/json"
	"errors"
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

func StatusCodeForError(err error, dupErr interface{}, customStatusCode int) int {
    if errors.As(err, &dupErr) {
        return customStatusCode
    }
    return http.StatusInternalServerError
}
