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
	w.Header().Set("Hx-Reswap", "main")
	w.Header().Set("Hx-Retarget", "innerHTML")
	utils.SendAuthPage(w, h.Config.BaseURL, "pages/register.html")
}

func (h *PagesHandler) LoginPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Hx-Reswap", "main")
	w.Header().Set("Hx-Retarget", "innerHTML")
	utils.SendAuthPage(w, h.Config.BaseURL, "pages/register.html")

}

func (h *PagesHandler) NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Hx-Retarget", "main")
	w.Header().Set("Hx-Reswap", "innerHTML")
	utils.SendHTMLDocumentResponse(w, nil, "pages/not_found.html")
}