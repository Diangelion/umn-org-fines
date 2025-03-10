package handlers

import (
	"gateway/utils"
	"log"
	"net/http"
)

func GetListOrganizations(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(string)
	if !ok {
		log.Println("GetListOrganizations | Context userId not found")
		utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName)
		return
	}
}

func GetSingleOrganization(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(string)
	if !ok {
		log.Println("GetSingleOrganization | Context userId not found")
		utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), fileName)
		return
	}
}

func CreateOrganization(w http.ResponseWriter, r *http.Request) {
}

func JoinOrganization(w http.ResponseWriter, r *http.Request) {
}