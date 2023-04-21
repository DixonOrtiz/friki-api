package jwtauth

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func GetExternalIDFromClaims(token *jwt.Token) (string, error) {
	claims, _ := token.Claims.(jwt.MapClaims)
	return fmt.Sprint(claims["external_id"]), nil
}
