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
	"backend/middleware"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.LoadConfig()

	// Connect to the database
	db, err := sql.Open("postgres", cfg.GetConnectionString())
	if err != nil {
		log.Println(err)
		return
	}

	// Verify database connection
	if err := db.Ping(); err != nil {
		log.Println(err)
		return
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
	authRouter.HandleFunc("/login", userHandler.Login).Methods("POST")

	// Wrap the router with your CORS middleware so that every request goes through it.
	handlerWithCORS := middleware.CORS(router)

	// Start the server
	serverURL := fmt.Sprintf("http://127.0.0.1:%s", cfg.HTTPPort)
	fmt.Println("Server running at:", serverURL)
	log.Fatal(http.ListenAndServe(":"+cfg.HTTPPort, handlerWithCORS))
}
