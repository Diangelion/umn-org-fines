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
		log.Println("Main | Connect to databaser error: ", err)
		return
	}

	// Verify database connection
	if err := db.Ping(); err != nil {
		log.Println("Main | Pinging databaser error: ", err)
		return
	}

	router := mux.NewRouter().StrictSlash(false)
	jwt := middleware.NewJWT(db, cfg)
	pagesHandler := handlers.NewPagesHandler(cfg)
	partialHandler := handlers.NewPartialHandler(cfg)

    // 1. Auth API Routes (should come first as they're most specific)
	// Public User Routes
    authRouter := router.PathPrefix("/auth").Subrouter()
    authRouter.HandleFunc("/register", handlers.RegisterUserHandler).Methods("POST")
    authRouter.HandleFunc("/login", handlers.LoginUserHandler).Methods("POST")
	// Protected User Routes
	protectedAuthRouter := authRouter.NewRoute().Subrouter()
	protectedAuthRouter.Use(jwt.ProtectedMiddleware)
	protectedAuthRouter.HandleFunc("/user", handlers.GetUserHandler).Methods("GET")
	protectedAuthRouter.HandleFunc("/edit", handlers.EditUserHandler).Methods("POST")
	// Protected Organization Routes
	orgRouter := router.PathPrefix("/organization").Subrouter()
	orgRouter.Use(jwt.ProtectedMiddleware)
	// orgRouter.HandleFunc("/get-all", handlers.GetListOrganizations).Methods("GET")
	// orgRouter.HandleFunc("/get", handlers.GetSingleOrganization).Methods("GET")
	orgRouter.HandleFunc("/create", handlers.CreateOrganizationHandler).Methods("POST")
	// orgRouter.HandleFunc("/join", handlers.JoinOrganization).Methods("POST")

    // 2. Protected Routes (pages/partial document requiring authentication)
	// Pages
    protectedRouter := router.NewRoute().Subrouter()
    protectedRouter.Use(jwt.ProtectedMiddleware)
    protectedRouter.HandleFunc("/home", pagesHandler.HomePage).Methods("GET")
    protectedRouter.HandleFunc("/profile", pagesHandler.ProfilePage).Methods("GET")
	// Partial
	partialRouter := protectedRouter.PathPrefix("/partial").Subrouter()
	partialRouter.HandleFunc("/sidebar-profile", partialHandler.SidebarProfilePartial).Methods("GET")
	partialRouter.HandleFunc("/sidebar-organization-list", partialHandler.SidebarOrganizationListPartial).Methods("GET")

    // 3. Public Routes (with redirect middleware)
    publicRouter := router.NewRoute().Subrouter()
    publicRouter.Use(jwt.PublicMiddleware)
    publicRouter.HandleFunc("/", pagesHandler.IndexPage).Methods("GET")
    publicRouter.HandleFunc("/login", pagesHandler.LoginPage).Methods("GET")
    publicRouter.HandleFunc("/register", pagesHandler.RegisterPage).Methods("GET")

    // 4. Not Found Handler (should be on main router)
    router.NotFoundHandler = http.HandlerFunc(pagesHandler.NotFoundPage)

	// Wrap all router with CORS middleware so that every request goes through it.
	handlerWithCORS := middleware.CORS(router)

	// Start the server
	serverURL := fmt.Sprintf("http://127.0.0.1:%s", cfg.HTTPPort)
	fmt.Println("Gateway running at", serverURL)
	log.Fatal(http.ListenAndServe(":"+cfg.HTTPPort, handlerWithCORS))

	// Defer
	defer db.Close()
}
