package utils

import (
	"net/http"

	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

// Parse form data and return an error if parsing fails
func ParseRequestBody(r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return err
	}
	return nil
}

// Decode form data into the provided struct and return an error if decoding fails
func DecodeRequestBody(r *http.Request, dest interface{}) error {
	if err := decoder.Decode(dest, r.PostForm); err != nil {
		return err
	}
	return nil
}
