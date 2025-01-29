package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"umn-org-fines/backend/config"
	"umn-org-fines/backend/internal/handlers"
	"umn-org-fines/backend/internal/repositories"
	"umn-org-fines/backend/internal/services"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Load .env file (Only for local development)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables.")
	}

	// Set the port (use environment variable or default to 8080)
	port := os.Getenv("HTTP_PORT")

	// Load configuration
	cfg := config.LoadConfig()

	// Connect to the database
	db, err := sql.Open("postgres", cfg.GetConnectionString())
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	// Verify database connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Database is not reachable: %v\n", err)
	}
	defer db.Close()

	// Initialize repository, service, and handler
	userRepo := repositories.NewUserRepository(db)
	userSvc := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userSvc)

	// Set up router using Gorilla Mux
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/register", userHandler.Register).Methods("POST")

	// Start the server
	serverURL := fmt.Sprintf("http://localhost:%s", port)
	fmt.Println("Server running at:", serverURL)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
