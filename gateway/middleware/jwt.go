package middleware

import (
	"database/sql"
	"errors"
	"gateway/config"
	"gateway/utils"
	"log"
	"net/http"
	"strings"

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

// isUserExists checks if the user exists in the database.
func (m *JWTMiddleware) isUserExists(userId string) bool {
	query := `SELECT 1 FROM user_credentials WHERE user_id = $1`
	var exists int
	err := m.DB.QueryRow(query, userId).Scan(&exists)
	return err == nil
}

func (m *JWTMiddleware) getJWTKey(whichToken string) []byte {
	returnVal := []byte(cfg.JWTAccessKey)
	if (whichToken == "refresh") {
		returnVal = []byte(cfg.JWTRefreshKey)
	}
	return returnVal
}

func (m *JWTMiddleware) ParseJWT(tokenValue string, tokenType string) (jwt.MapClaims, error) {
	// Parse token value
	parsedToken, err :=  jwt.Parse(tokenValue, func(t *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC.
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Printf("ParseJWT | Unexpected signing method: %v", t.Header["alg"])
			return nil, errors.New("Invalid signing method")
		}
		return m.getJWTKey(tokenType), nil
	})

	// Error handling if error occured / invalid parsed token / token expired
	if err != nil || !parsedToken.Valid {
		log.Println("ParseJWT | Parse token error: ", err)
		return nil, errors.New("JWT is expired or invalid")
	}
	
	// Extract claims from the token
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		log.Printf("ParseJWT | Invalid token claims\n")
		return nil, errors.New("Invalid token claims")
	}

	return claims, nil
}

// verifyToken checks for access_token first, then refresh_token if access_token is invalid.
func (m *JWTMiddleware) verifyToken(w http.ResponseWriter, r *http.Request) (string, error) {
	// Get access token first first
	accessToken := r.Header.Get("Authorization")
	if accessToken != "" {
		if parts := strings.Split(accessToken, " "); len(parts) == 2 && parts[0] == "Bearer" {
			if claims, err := m.ParseJWT(parts[1], "access"); err == nil {
				if userId, ok := claims["user_id"].(string); ok {
					return userId, nil
				}
			}
		}
	}

	// If access token is invalid, check refresh token
	refreshToken := r.Header.Get("X-Refresh-Token")
	if refreshToken == "" {
		log.Printf("verifyToken | Missing refresh token\n")
		return "", errors.New("Missing refresh token")
	}

	claims, err := m.ParseJWT(refreshToken, "refresh")
	if err != nil {
		log.Println("verifyToken | Parse JWT error: ", err)
		return "", err
	}
	
	userId, ok := claims["user_id"].(string)
	if !ok {
		log.Printf("verifyToken | Invalid token claims\n")
		return "", errors.New("Invalid token claims")
	}
	
	// Generate a new access token
	newAccessToken, err := utils.GenerateAccessToken(userId)
	if err != nil {
		log.Println("verifyToken | Generate access token error: ", err)
		return "", errors.New("Unable to generate access token")
	}

	w.Header().Set("Authorization", newAccessToken)
	w.Header().Set("HX-Trigger", "refreshAccessToken()")

	return userId, nil
}

// JWTMiddleware verifies the JWT token in the Authorization header.
func (m *JWTMiddleware) ProtectedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId, err := m.verifyToken(w, r)
		if err != nil {
			log.Println("ProtectedMiddleware | Verify token error: ", err)
			w.Header().Set("HX-Redirect", "/login")
			w.WriteHeader(http.StatusAccepted)
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


// Redirect verifies the JWT token in the Authorization header.
// If exist then redirect to homepage, otherwise proceed unauthorized page like login, register, etc.
func (m *JWTMiddleware) PublicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId, err := m.verifyToken(w, r)
		if err == nil && m.isUserExists(userId) {
			log.Printf("PublicMiddleware | Session exist\n")
			w.Header().Set("HX-Redirect", "/home")
			w.WriteHeader(http.StatusAccepted)
			return
		}

		next.ServeHTTP(w, r)
	})
}