package main

import (
	"fmt"
	"gateway/config"
	"gateway/internal/handlers"
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
	router.HandleFunc("/api/register", handlers.RegisterUser).Methods("POST")

	// Start the server
	serverURL := fmt.Sprintf("http://localhost:%s", cfg.HTTPPort)
	fmt.Println("Gateway running at", serverURL)
	log.Fatal(http.ListenAndServe(":"+cfg.HTTPPort, router))
}
