package middleware

import (
	"net/http"
)

// CORS middleware sets CORS headers on the response.
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:3334")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")

		// If it's a preflight OPTIONS request, respond with OK directly.
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Otherwise, proceed with the next handler.
		next.ServeHTTP(w, r)
	})
}
