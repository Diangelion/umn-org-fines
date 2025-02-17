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


type JWT struct {
    db *sql.DB
}

func NewJWT(db *sql.DB) *JWT {
    return &JWT{db}
}

var cfg = config.LoadConfig()

func (j *JWT) getJWTKey(whichToken string) (interface{}) {
	returnVal := cfg.JWTAccessKey
	if (whichToken == "refresh") {
		returnVal = cfg.JWTRefreshKey
	}
	return returnVal
}

func (j *JWT) ParseJWT(tokenValue string, tokenType string) (jwt.MapClaims, error) {
	// Parse token value
	parsedToken, err :=  jwt.Parse(tokenValue, func(t *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC.
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Println(fmt.Printf("Unexpected signing method: %v", t.Header["alg"]))
			return nil, errors.New("Invalid signing method")
		}
		return j.getJWTKey(tokenType), nil
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
func (j *JWT) JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId, err := j.verifyToken(w, r)
		if err != nil {
			log.Println(err)
			w.Header().Set("HX-Reswap", "main")
			w.Header().Set("HX-Retarget", "outerHTML")
			w.WriteHeader(http.StatusOK)
			utils.SendHTMLDocumentResponse(w, nil, "pages/login.html", http.StatusAccepted)
			return
		}

		// Validate user existence in DB
		if !j.isUserExists(userId) {
			utils.SendAlert(w, "Error", "Unauthorized", "alert.html", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// verifyToken checks for access_token first, then refresh_token if access_token is invalid.
func (j *JWT) verifyToken(w http.ResponseWriter, r *http.Request) (string, error) {
	// Try access_token first
	if accessToken, err := r.Cookie("access_token"); err == nil {
		if claims, err := j.ParseJWT(accessToken.Value, "access"); err == nil {
			if userId, ok := claims["user_id"].(string); ok {
				return userId, nil
			}
		}
	}

	// If access_token is invalid, check refresh_token
	refreshToken, err := r.Cookie("refresh_token")
	if err != nil {
		return "", err
	}

	claims, err := j.ParseJWT(refreshToken.Value, "refresh")
	if err != nil {
		log.Println(err)
		return "", err
	}

	userId, ok := claims["user_id"].(string)
	if !ok {
		log.Printf("Invalid token claims")
		return "", errors.New("Invalid token claims")
	}

	// Generate a new access token
	newAccessToken, err := utils.GenerateAccessToken(userId)
	if err != nil {
		log.Println(err)
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
func (j *JWT) isUserExists(userId string) bool {
	query := `SELECT 1 FROM user_credentials WHERE user_id = $1`
	var exists int
	err := j.db.QueryRow(query, userId).Scan(&exists)
	return err == nil
}

