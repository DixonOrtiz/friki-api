package oauthrepo

import (
	oauthinfra "frikiapi/src/infraestructure/oauth"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateLoginURL(t *testing.T) {
	config := new(oauthinfra.MockOAuthConfig)
	config.On("AuthCodeURL").Return(testLoginURL)
	authRepository := MakeOAuthRepository(config)

	URL := authRepository.GenerateLoginURL()

	assert.Equal(t, testLoginURL, URL)
}
