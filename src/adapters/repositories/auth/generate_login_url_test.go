package authrepo

import (
	authinfra "frikiapi/src/infraestructure/auth"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateLoginURL(t *testing.T) {
	config := new(authinfra.MockAuthConfig)
	config.On("AuthCodeURL").Return(testLoginURL)
	authRepository := MakeAuthRepository(config)

	URL := authRepository.GenerateLoginURL()

	assert.Equal(t, testLoginURL, URL)
}
