package handlers

import (
	"backend/internal/models"
	"backend/internal/services"
	"backend/utils"
	"log"
	"net/http"
)

type OrganizationHandler struct {
    service *services.OrganizationService
}

func NewOrganizationHandler(service *services.OrganizationService) *OrganizationHandler {
    return &OrganizationHandler{service}
}

func (h *OrganizationHandler) GetListOrganizationHandler(w http.ResponseWriter, r *http.Request) {
    userId := r.Header.Get("Authorization")

    // Perform registration
    listOrganization, err := h.service.GetListOrganizationService(userId);
    if err != nil {
        log.Println("GetListOrganizationHandler | Get list organization service error: ", err)
        utils.SendJSONResponse(w, http.StatusInternalServerError, err.Error(), nil)
        return
    }

    utils.SendJSONResponse(w, http.StatusOK, "List organization has been successfully gained.", listOrganization)
}

func (h *OrganizationHandler) CreateOrganizationHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("Authorization")
    var org *models.CreateOrganization

    // Assign request body to user for validation
    if err := utils.DecodeRequestBody(r, org); err != nil {
        log.Println("CreateOrganizationHandler | Decode request error: ", err)
        utils.SendJSONResponse(w, http.StatusBadRequest, "Invalid JSON payload.", nil)
        return
    }

    // Perform registration
    if err := h.service.CreateOrganizationService(org, userId); err != nil {
        log.Println("CreateOrganizationHandler | Create organization service error: ", err)

        // Handle duplicate email specifically, differentiate with response status code
        var dupErr *models.DuplicateEmailError
        statusCode := utils.StatusCodeForError(err, dupErr, http.StatusConflict)
        utils.SendJSONResponse(w, statusCode, err.Error(), nil)
        return
    }

    utils.SendJSONResponse(w, http.StatusCreated, "Organization has been successfully created.", org)
}