package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config struct holds all the environment variables
type Config struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	HTTPPort   string
}

// LoadConfig reads from environment variables or .env file
func LoadConfig() *Config {
	// Load .env file (Only in local development)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables.")
	}

	return &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
		HTTPPort:   os.Getenv("HTTP_PORT"),
	}
}

// GetConnectionString returns a PostgreSQL connection string
func (c *Config) GetConnectionString() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		c.DBHost, c.DBUser, c.DBPassword, c.DBName, c.DBPort,
	)
}
