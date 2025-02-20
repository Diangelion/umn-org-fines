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
	http.Redirect(w, r, "/auth/is-logged-in", http.StatusFound)
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

func (h *PagesHandler) HomePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("HX-Reswap", "innerHTML")
	w.Header().Set("HX-Retarget", "main")
	utils.SendAuthPage(w, h.Config.BaseURL, "pages/home.html")

}

func (h *PagesHandler) NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("HX-Reswap", "innerHTML")
	w.Header().Set("HX-Retarget", "main")
	utils.SendHTMLDocumentResponse(w, nil, "pages/not_found.html")
}