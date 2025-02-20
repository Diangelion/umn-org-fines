package middleware

import (
	"net/http"
)

// CORS middleware sets CORS headers on the response.
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set(
			"Access-Control-Allow-Headers", 
			"Content-Type, Authorization, HX-Request, HX-Current-Url, HX-Trigger, HX-Target",
		)
		w.Header().Set("Access-Control-Expose-Headers", "HX-Redirect, HX-Reswap, HX-Retarget") // ✅ Allow frontend to read it
		w.Header().Set("Content-Type", "text/html")

		// If it's a preflight OPTIONS request, respond with OK directly.
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Otherwise, proceed with the next handler.
		next.ServeHTTP(w, r)
	})
}
