package jwtauth

import (
	"os"
	gotime "time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     gotime.Now().Add(gotime.Hour * 24 * 7).Unix(),
		"user_id": userID,
	})
	return token.SignedString([]byte(os.Getenv("JWT_KEY")))
}
