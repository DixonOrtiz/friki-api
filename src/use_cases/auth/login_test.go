package authusecases

import (
	"errors"
	"fmt"
	authrepo "frikiapi/src/adapters/repositories/auth"
	userrepo "frikiapi/src/adapters/repositories/user"
	"frikiapi/src/entities"
	jwtauth "frikiapi/src/utils/jwt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessfulLogin(t *testing.T) {
	authRepository := new(authrepo.MockAuthRepository)
	authRepository.On("GenerateExternalToken").Return(testExternalToken, nil)
	userRepository := new(userrepo.MockUserRepository)
	userRepository.On("GetExternalUserByToken").Return(testExternalUser, nil)
	userRepository.On("GetByExternalID").Return(testUser, nil)
	authUseCases := MakeAuthUseCases(authRepository, userRepository)

	token, err := authUseCases.Login("test_code")

	assert.NotEmpty(t, token)
	assert.Nil(t, err)
}

func TestLoginWithErrorInExternalTokenGeneration(t *testing.T) {
	testError := "there was an error generating external token"
	authRepository := new(authrepo.MockAuthRepository)
	authRepository.On("GenerateExternalToken").Return("", errors.New(testError))
	userRepository := new(userrepo.MockUserRepository)
	authUseCases := MakeAuthUseCases(authRepository, userRepository)

	token, err := authUseCases.Login("test_code")

	assert.Empty(t, token)
	assert.EqualError(t, err, fmt.Sprintf("internal: %s", testError))
}

func TestLoginWithErrorGettingExternalUserByToken(t *testing.T) {
	testError := "there was an error obtaining external user"
	authRepository := new(authrepo.MockAuthRepository)
	authRepository.On("GenerateExternalToken").Return("", nil)
	userRepository := new(userrepo.MockUserRepository)
	userRepository.On("GetExternalUserByToken").Return(entities.User{}, errors.New(testError))

	authUseCases := MakeAuthUseCases(authRepository, userRepository)

	token, err := authUseCases.Login("test_code")

	assert.Empty(t, token)
	assert.EqualError(t, err, fmt.Sprintf("internal: %s", testError))
}

func TestLoginWithNewUserRegistration(t *testing.T) {
	authRepository := new(authrepo.MockAuthRepository)
	authRepository.On("GenerateExternalToken").Return(testExternalToken, nil)
	userRepository := new(userrepo.MockUserRepository)
	userRepository.On("GetExternalUserByToken").Return(testExternalUser, nil)
	authUseCases := MakeAuthUseCases(authRepository, userRepository)

	tokenStr, err := authUseCases.Login("test_code")

	externalID, _ := jwtauth.GetClaimsFromTestToken(tokenStr)

	assert.Equal(t, expectedExternalID, externalID)
	assert.Nil(t, err)
}

func TestLoginWithErrorInJWTGeneration(t *testing.T) {
	testError := "there was an error in jwt generation"
	authRepository := new(authrepo.MockAuthRepository)
	authRepository.On("GenerateExternalToken").Return(testExternalToken, nil)
	userRepository := new(userrepo.MockUserRepository)
	userRepository.On("GetExternalUserByToken").Return(testExternalUser, nil)

	authUseCases := MakeAuthUseCases(authRepository, userRepository)
	patchGenerateJWT(testError)

	token, err := authUseCases.Login("test_code")

	assert.Empty(t, token)
	assert.EqualError(t, err, fmt.Sprintf("internal: %s", testError))
}
