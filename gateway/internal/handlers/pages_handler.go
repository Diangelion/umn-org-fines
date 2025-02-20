package handlers

import (
	"gateway/config"
	"gateway/utils"
	"net/http"
)

type PagesHandler struct {
	Config *config.Config
}

func NewPagesHandler(cfg *config.Config) *PagesHandler {
    return &PagesHandler{Config: cfg}
}

func (h *PagesHandler) IndexPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("HX-Redirect", "/login")
	w.WriteHeader(http.StatusOK)
}

func (h *PagesHandler) RegisterPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("HX-Reswap", "innerHTML")
	w.Header().Set("HX-Retarget", "main")
	utils.SendAuthPage(w, h.Config.BaseURL, "pages/register.html")
}

func (h *PagesHandler) LoginPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("HX-Reswap", "innerHTML")
	w.Header().Set("HX-Retarget", "main")
	utils.SendAuthPage(w, h.Config.BaseURL, "pages/login.html")

}

func (h *PagesHandler) NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("HX-Reswap", "innerHTML")
	w.Header().Set("HX-Retarget", "main")
	utils.SendHTMLDocumentResponse(w, nil, "pages/not_found.html")
}