package utils

import (
	"encoding/json"
	"net/http"
)

func DecodeRequestBody(r *http.Request, user any) error {
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return err
	}
	return nil
}