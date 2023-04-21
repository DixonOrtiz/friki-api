package jwtauth

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetExternalIDFromClaims(t *testing.T) {
	os.Setenv("JWT_KEY", testJWTKey)
	defer os.Clearenv()
	token := getTestToken(externalIDInput)

	externalID, err := GetExternalIDFromClaims(token)

	assert.Equal(t, expectedExternalID, externalID)
	assert.Nil(t, err)
}
