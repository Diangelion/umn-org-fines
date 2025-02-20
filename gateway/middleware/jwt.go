package middleware

import (
	"database/sql"
	"errors"
	"fmt"
	"gateway/config"
	"gateway/utils"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


type JWTMiddleware struct {
    DB *sql.DB
	Config *config.Config
}

func NewJWT(db *sql.DB, cfg *config.Config) *JWTMiddleware {
    return &JWTMiddleware{DB: db, Config: cfg}
}

var cfg = config.LoadConfig()

func (m *JWTMiddleware) getJWTKey(whichToken string) (interface{}) {
	returnVal := cfg.JWTAccessKey
	if (whichToken == "refresh") {
		returnVal = cfg.JWTRefreshKey
	}
	return returnVal
}

func (m *JWTMiddleware) ParseJWT(tokenValue string, tokenType string) (jwt.MapClaims, error) {
	// Parse token value
	parsedToken, err :=  jwt.Parse(tokenValue, func(t *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC.
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Println(fmt.Printf("Unexpected signing method: %v", t.Header["alg"]))
			return nil, errors.New("Invalid signing method")
		}
		return m.getJWTKey(tokenType), nil
	})

	// Error handling if error occured / invalid parsed token
	if err != nil || !parsedToken.Valid {
		log.Println(err.Error())
		return nil, errors.New("JWT is expired or invalid")
	}
	
	// Extract claims from the token
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Invalid token claims")
	}

	return claims, nil
}

// JWTMiddleware verifies the JWT token in the Authorization header.
func (m *JWTMiddleware) JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId, err := m.verifyToken(w, r)
		if err != nil {
			log.Println("JWTMiddleware | Verify token error: ", err)
			w.Header().Set("HX-Redirect", "/login")
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Validate user existence in DB
		if !m.isUserExists(userId) {
			utils.SendAlert(w, "Error", "Invalid user.", "alert.html")
			return
		}

		next.ServeHTTP(w, r)
	})
}

// verifyToken checks for access_token first, then refresh_token if access_token is invalid.
func (m *JWTMiddleware) verifyToken(w http.ResponseWriter, r *http.Request) (string, error) {
	// Try access_token first
	if accessToken, err := r.Cookie("access_token"); err == nil {
		if claims, err := m.ParseJWT(accessToken.Value, "access"); err == nil {
			if userId, ok := claims["user_id"].(string); ok {
				return userId, nil
			}
		}
	}

	// If access_token is invalid, check refresh_token
	refreshToken, err := r.Cookie("refresh_token")
	if err != nil {
		log.Println("verifyToken | Cookie error: ", err)
		return "", err
	}

	claims, err := m.ParseJWT(refreshToken.Value, "refresh")
	if err != nil {
		log.Println("verifyToken | Parse JWT error: ", err)
		return "", err
	}
	
	userId, ok := claims["user_id"].(string)
	if !ok {
		log.Printf("verifyToken | Invalid token claims")
		return "", errors.New("Invalid token claims")
	}
	
	// Generate a new access token
	newAccessToken, err := utils.GenerateAccessToken(userId)
	if err != nil {
		log.Println("verifyToken | Generate access token error: ", err)
		return "", errors.New("Unable to generate access token")
	}

	// Set new access_token in cookies
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    newAccessToken,
		HttpOnly: true,
		Secure:   false, // Change to true if using HTTPS
		Path:     "/",
		Expires:  time.Now().Add(15 * time.Minute),
	})

	return userId, nil
}

// isUserExists checks if the user exists in the database.
func (m *JWTMiddleware) isUserExists(userId string) bool {
	query := `SELECT 1 FROM user_credentials WHERE user_id = $1`
	var exists int
	err := m.DB.QueryRow(query, userId).Scan(&exists)
	return err == nil
}

