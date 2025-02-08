package main

import (
	"fmt"
	"gateway/config"
	"gateway/internal/handlers"
	"gateway/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	cfg := config.LoadConfig()
	router := mux.NewRouter()

	// // Apply middleware
	// router.Use(middlewares.LoggingMiddleware)

	// Define routes
	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/register", handlers.RegisterUser).Methods("POST")
	authRouter.HandleFunc("/login", handlers.LoginUser).Methods("POST")

	// Wrap the router with your CORS middleware so that every request goes through it.
	handlerWithCORS := middleware.CORS(router)

	// Start the server
	serverURL := fmt.Sprintf("http://localhost:%s", cfg.HTTPPort)
	fmt.Println("Gateway running at", serverURL)
	log.Fatal(http.ListenAndServe(":"+cfg.HTTPPort, handlerWithCORS))
}
