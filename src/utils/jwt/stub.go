package jwtauth

import (
	"github.com/golang-jwt/jwt/v5"
)

var testJWTKey = "secret_key"
var IDInput = "12345"
var expectedID = "12345"
var expectedExp = 1.678717665e+09
var testTokenInput = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Nzg3MTc2NjUsImV4dGVybmFsX2lkIjoiMTIzNCJ9.qgNpbWlVTjYQRbAT03ojLJZocy4z6RVuAppzRiQihBw"

func GetClaimsFromTestToken(tokenStr string) (userID string, exp float64) {
	token, _ := ParseJWT(tokenStr)
	claims, _ := token.Claims.(jwt.MapClaims)
	return claims["user_id"].(string), claims["exp"].(float64)
}

func getTestToken(userID string) *jwt.Token {
	tokenStr, _ := GenerateJWT(userID)
	token, _ := ParseJWT(tokenStr)
	return token
}
