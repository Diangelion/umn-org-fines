package main

import (
	"database/sql"
	"fmt"
	"log"
	"umn-org-fines/config"

	_ "github.com/lib/pq"
)

func main() {
	cfg := config.LoadConfig()

	// Connect to the database
	db, err := sql.Open("postgres", cfg.GetConnectionString())
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Ping to database failed: %v\n", err)
	}

	fmt.Println("Connected to the database successfully!")
}
