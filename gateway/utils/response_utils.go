package utils

import (
	"gateway/internal/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func getAlert(title string, errMsg string) interface{} {
	return models.Alert{
		Title: title,
		Message: errMsg,
	}
}

func SendHTMLDocumentResponse(w http.ResponseWriter, data interface{}, fileName string, statusCode int) {
	// Locate the template file
	templatePath := filepath.Join("templates", fileName)
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Println("Template parse error:", err)
		return
	}

	// Custom httpStatus code (not always 200)
	// Set status code BEFORE writing content
	w.WriteHeader(statusCode)

	// Execute the template, writing the output to the ResponseWriter.
	if err := tmpl.Execute(w, data); err != nil {
		log.Println("Template execution error:", err)
		return
	}
}

func SendAlert(w http.ResponseWriter, alertTitle string, errMsg string, fileName string, statusCode int) {
	document := getAlert(alertTitle, errMsg)
	SendHTMLDocumentResponse(w, document, fileName, statusCode)
}