package jwtauth

import (
	"frikiapi/src/utils"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateJWT(t *testing.T) {
	os.Setenv("JWT_KEY", testJWTKey)
	utils.PatchTimeNow()

	tokenStr, err := GenerateJWT(externalIDInput)
	externalID, exp := GetClaimsFromTestToken(tokenStr)

	assert.Equal(t, expectedExternalID, externalID)
	assert.Equal(t, expectedExp, exp)
	assert.Nil(t, err)
}
