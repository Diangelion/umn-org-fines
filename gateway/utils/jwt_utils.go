package utils

import (
	"gateway/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var c = config.LoadConfig()

// GenerateAccessToken generates a JWT access token with a short expiration (e.g., 15 minutes)
func GenerateAccessToken(userId string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Convert c.JWTAccessKey to []byte
	return token.SignedString([]byte(c.JWTAccessKey))
}

// GenerateRefreshToken generates a JWT refresh token with a longer expiration (e.g., 1 day)
func GenerateRefreshToken(userId string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Convert c.JWTRefreshKey to []byte
	return token.SignedString([]byte(c.JWTRefreshKey))
}
