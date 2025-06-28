package helper

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("ACUMALAKAXIXIXIXI")

func GenerateJWT(userId int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(24 * time.Minute).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
