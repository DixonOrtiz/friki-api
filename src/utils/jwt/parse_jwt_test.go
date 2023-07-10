package jwtauth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseJWT(t *testing.T) {
	tokenStr, _ := GenerateJWT(IDInput)

	token, err := ParseJWT(tokenStr)

	assert.True(t, token.Valid)
	assert.Nil(t, err)
}

func TestParseInvalidJWT(t *testing.T) {
	token, err := ParseJWT(testTokenInput)

	assert.Nil(t, token)
	assert.EqualError(t, err, "token signature is invalid: signature is invalid")
}
