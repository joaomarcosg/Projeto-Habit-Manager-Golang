package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	JWT_SECRET_KEY = "JWT_SECRET_KEY"
)

func GenerateToken(userID, email string) (string, error) {

	secret := os.Getenv(JWT_SECRET_KEY)

	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	return "", nil
}
