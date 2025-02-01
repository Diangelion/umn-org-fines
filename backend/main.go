package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"backend/config"
	"backend/internal/handlers"
	"backend/internal/repositories"
	"backend/internal/services"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.LoadConfig()

	// Connect to the database
	db, err := sql.Open("postgres", cfg.GetConnectionString())
	if err != nil {
		log.Fatalf("Failed: unable to connect to database: %v\n", err)
	}

	// Verify database connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed: database is not reachable: %v\n", err)
	}
	defer db.Close()

	// Set up router using Gorilla Mux
	router := mux.NewRouter()

	// User
	userRepo := repositories.NewUserRepository(db)
	userSvc := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userSvc)

	authRouter := router.PathPrefix("/user").Subrouter()
	authRouter.HandleFunc("/register", userHandler.Register).Methods("POST")

	// Start the server
	serverURL := fmt.Sprintf("http://localhost:%s", cfg.HTTPPort)
	fmt.Println("Server running at:", serverURL)
	log.Fatal(http.ListenAndServe(":"+cfg.HTTPPort, router))
}
