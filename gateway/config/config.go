package config

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

// Config struct holds all the environment variables
type Config struct {
	HTTPPort      string
	BaseURL       string
	BackendURL    string
	DBHost        string
	DBUser        string
	DBPassword    string
	DBName        string
	DBPort        string
	JWTAccessKey  string
	JWTRefreshKey string
}

var (
	cfg  *Config       // Declare cfg at the package level
	once sync.Once     // Declare once at the package level
)

// LoadConfig reads from environment variables or .env file
func LoadConfig() *Config {
	once.Do(func() {
		// Load .env file (Only in local development)
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found, using system environment variables.")
		}

		// Initialize cfg inside the once.Do block
		cfg = &Config{
			HTTPPort:      os.Getenv("PORT"),
			BaseURL:       os.Getenv("BASE_URL"),
			BackendURL:    os.Getenv("BACKEND_URL"),
			DBHost:        os.Getenv("DB_HOST"),
			DBUser:        os.Getenv("DB_USER"),
			DBPassword:    os.Getenv("DB_PASSWORD"),
			DBName:        os.Getenv("DB_NAME"),
			DBPort:        os.Getenv("DB_PORT"),
			JWTAccessKey:  os.Getenv("JWT_ACCESS_KEY"),
			JWTRefreshKey: os.Getenv("JWT_REFRESH_KEY"),
		}

		log.Println("Configuration loaded")
	})

	return cfg // Return the package-level cfg
}

// GetConnectionString returns a PostgreSQL connection string
func (c *Config) GetConnectionString() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		c.DBHost, c.DBUser, c.DBPassword, c.DBName, c.DBPort,
	)
}