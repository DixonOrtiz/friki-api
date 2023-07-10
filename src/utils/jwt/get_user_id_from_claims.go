package jwtauth

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func GetUserIDFromClaims(token *jwt.Token) (string, error) {
	claims, _ := token.Claims.(jwt.MapClaims)
	return fmt.Sprint(claims["user_id"]), nil
}
