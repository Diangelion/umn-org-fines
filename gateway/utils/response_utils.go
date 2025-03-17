package utils

import (
	"gateway/internal/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func SendHTMLDocumentResponse(w http.ResponseWriter, data any, fileName string) {
	// Locate the template file
	templatePath := filepath.Join("templates", fileName)
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Println("Template parse error:", err)
		return
	}

	// Set status code BEFORE writing content
	// Set always 200 to make hx-swap, hx-target, etc. running
	w.WriteHeader(http.StatusAccepted)

	// Execute the template, writing the output to the ResponseWriter.
	if err := tmpl.Execute(w, data); err != nil {
		log.Println("Template execution error:", err)
		return
	}
}

// Getter
func getAlert(title string, errMsg string) any {
	return models.Alert{
		Title: title,
		Message: errMsg,
	}
}

func GetAuthPage(baseURL string) any {
	return models.AuthPage{BaseURL: baseURL}
}

// Setter
func SendAlert(w http.ResponseWriter, alertTitle string, errMsg string, fileName string) {
	document := getAlert(alertTitle, errMsg)
	SendHTMLDocumentResponse(w, document, fileName)
}


func SendAuthPage(w http.ResponseWriter, document any, fileName string) {
	SendHTMLDocumentResponse(w, document, fileName)
}

func SendAuthPartial(w http.ResponseWriter, document any, fileName string) {
	SendHTMLDocumentResponse(w, document, fileName)
}