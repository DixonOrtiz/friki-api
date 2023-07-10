package jwtauth

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserIDFromClaims(t *testing.T) {
	os.Setenv("JWT_KEY", testJWTKey)
	defer os.Clearenv()
	token := getTestToken(IDInput)

	externalID, err := GetUserIDFromClaims(token)

	assert.Equal(t, expectedID, externalID)
	assert.Nil(t, err)
}
