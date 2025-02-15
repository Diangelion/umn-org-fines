package main

import (
	"database/sql"
	"fmt"
	"gateway/config"
	"gateway/internal/handlers"
	"gateway/middleware"
	"log"
	"net/http"

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
	
	router := mux.NewRouter()
	
	// Routes
	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/register", handlers.RegisterUser).Methods("POST")
	authRouter.HandleFunc("/login", handlers.LoginUser).Methods("POST")
	
	newJWT := middleware.NewJWT(db)
	protectedAuthRouter := authRouter.NewRoute().Subrouter()
	protectedAuthRouter.Use(newJWT.JWTMiddleware)
	protectedAuthRouter.HandleFunc("/is-logged-in", handlers.IsLoggedIn).Methods("GET")
	
	// Wrap all router with CORS middleware so that every request goes through it.
	handlerWithCORS := middleware.CORS(router)
	
	// Start the server
	serverURL := fmt.Sprintf("http://localhost:%s", cfg.HTTPPort)
	fmt.Println("Gateway running at", serverURL)
	log.Fatal(http.ListenAndServe(":"+cfg.HTTPPort, handlerWithCORS))

	// Defer
	defer db.Close()
}
