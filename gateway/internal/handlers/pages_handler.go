package handlers

import (
	"gateway/utils"
	"net/http"
)

func IndexPage(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/auth/is-logged-in", http.StatusFound)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("HX-Retarget", "main")
	w.Header().Set("HX-Reswap", "innerHTML")
	utils.SendHTMLDocumentResponse(w, nil, "pages/not_found.html", http.StatusOK)
}