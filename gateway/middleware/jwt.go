package middleware

import (
	"database/sql"
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

func (j *JWT) parseJWT(tokenValue string, tokenType string) (jwt.MapClaims, error) {
	// Parse token value
	parsedToken, err :=  jwt.Parse(tokenValue, func(t *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC.
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return j.getJWTKey(tokenType), nil
	})

	// Error handling if error occured / invalid parsed token
	if err != nil || !parsedToken.Valid {
		log.Println(err.Error())
		return nil, fmt.Errorf("jwt token expired or invalid")
	}
	
	// Extract claims from the token
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("token claims invalid")
	}

	return claims, nil
}

// JWTMiddleware verifies the JWT token in the Authorization header.
func (j *JWT) JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the access_token inside cookie header
		accessToken, err := r.Cookie("access_token")
		if err == nil {
			_, err := j.parseJWT(accessToken.Value, "access")
			if err == nil {
				next.ServeHTTP(w, r)
				return
			}
		}
		
		// Get the refresh_token inside cookie header
		refreshToken, err := r.Cookie("refresh_token")
		if err != nil {
			utils.SendAlert(w, "Error", "Unauthorized", "alert.html", http.StatusUnauthorized)
			return
		}
		
		// Verify refresh token
		claims, err := j.parseJWT(refreshToken.Value, "refresh")
		if err != nil {
			utils.SendAlert(w, "Error", "Unauthorized", "alert.html", http.StatusUnauthorized)
			return
		}

		// Get the email from the token claims
		email, ok := claims["email"].(string)
		if !ok {
			utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), "alert.html", http.StatusInternalServerError)
			return
		}

		// Generate a new access token
		newAccessToken, err := utils.GenerateAccessToken(email)
		if err != nil {
			utils.SendAlert(w, "Error", utils.GetGeneralErrorMessage(), "alert.html", http.StatusInternalServerError)
			return
		}

		// Set the new access token in the response cookies
		http.SetCookie(w, &http.Cookie{
			Name:     "access_token",
			Value:    newAccessToken,
			HttpOnly: true,
			Secure:   false, // Set to true if using HTTPS
			Path:     "/",
			Expires:  time.Now().Add(15 * time.Minute),
		})

		next.ServeHTTP(w, r)
	})
}
