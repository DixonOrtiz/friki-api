package jwtauth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(externalID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":         time.Now().Add(time.Hour * 24 * 7).Unix(),
		"external_id": externalID,
	})
	return token.SignedString([]byte(os.Getenv("JWT_KEY")))
}
