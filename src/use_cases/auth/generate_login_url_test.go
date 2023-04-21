package authusecases

import (
	authrepo "frikiapi/src/adapters/repositories/auth"
	userrepo "frikiapi/src/adapters/repositories/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateLoginURL(t *testing.T) {
	mockAuthRepository := new(authrepo.MockAuthRepository)
	mockAuthRepository.On("GenerateLoginURL").Return(testURLInput)
	mockUserRepository := new(userrepo.MockUserRepository)
	authUseCases := MakeAuthUseCases(mockAuthRepository, mockUserRepository)

	URL := authUseCases.GenerateLoginURL()

	assert.Equal(t, expectedTestURL, URL)
}
