package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config struct holds all the environment variables
type Config struct {
	HTTPPort   string
	BackendURL string
}

// LoadConfig reads from environment variables or .env file
func LoadConfig() *Config {
	// Load .env file (Only in local development)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables.")
	}

	return &Config{
		HTTPPort:   os.Getenv("PORT"),
		BackendURL:   os.Getenv("BACKEND_URL"),
	}
}
