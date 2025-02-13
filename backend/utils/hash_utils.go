package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	if hashedBytes, errHash := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); errHash != nil {
		return "", errHash
	} else {
		return string(hashedBytes), nil
	}
}


func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}