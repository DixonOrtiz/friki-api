package authrepo

import (
	"errors"
	authinfra "frikiapi/src/infraestructure/auth"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateExternalToken(t *testing.T) {
	testToken := getTestToken()
	config := new(authinfra.MockAuthConfig)
	config.On("Exchange").Return(testToken, nil)
	authRepository := MakeAuthRepository(config)

	token, err := authRepository.GenerateExternalToken("test_code")

	assert.Equal(t, expectedExternalToken, token)
	assert.Nil(t, err)
}

func TestGenerateExternalTokenWithError(t *testing.T) {
	testError := "there was an error exchanging google user code"
	config := new(authinfra.MockAuthConfig)
	config.On("Exchange").Return(emptyToken, errors.New(testError))
	authRepository := MakeAuthRepository(config)

	token, err := authRepository.GenerateExternalToken("test_code")

	assert.Equal(t, "", token)
	assert.EqualError(t, err, testError)
}
