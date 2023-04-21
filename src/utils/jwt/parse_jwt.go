package jwtauth

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func ParseJWT(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}
