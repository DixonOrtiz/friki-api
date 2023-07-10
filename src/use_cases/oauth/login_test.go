package authusecases

import (
	goerrors "errors"
	oauthrepository "frikiapi/src/adapters/repositories/oauth"
	googleuserrepo "frikiapi/src/adapters/repositories/users/google"
	"frikiapi/src/entities"
	userusecases "frikiapi/src/use_cases/users"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginWithErrorGeneratingExternalToken(t *testing.T) {
	oAuthRepository := new(oauthrepository.MockOAuthRepository)
	oAuthRepository.On("GenerateExternalToken").Return(
		"",
		goerrors.New("there was an error generating external token"),
	)

	googleUserRepository := new(googleuserrepo.MockGoogleUserRepository)
	userUseCases := new(userusecases.MockUserUseCases)

	oAuthUseCases := MakeOAuthUseCases(oAuthRepository, googleUserRepository, userUseCases)

	token, created, err := oAuthUseCases.Login("test_code")

	assert.Equal(t, "", token)
	assert.False(t, created)
	assert.ErrorContains(t, err, "internal: there was an error generating external token")
}

func TestLoginWithErrorGetingExternalUser(t *testing.T) {
	oAuthRepository := new(oauthrepository.MockOAuthRepository)
	oAuthRepository.On("GenerateExternalToken").Return(
		"test_external_token",
		nil,
	)
	externalUserRepository := new(googleuserrepo.MockGoogleUserRepository)
	externalUserRepository.On("GetByToken").Return(
		entities.User{},
		goerrors.New("there was an error getting user by token"),
	)
	userUseCases := new(userusecases.MockUserUseCases)

	oAuthUseCases := MakeOAuthUseCases(oAuthRepository, externalUserRepository, userUseCases)

	token, created, err := oAuthUseCases.Login("test_code")

	assert.Equal(t, "", token)
	assert.False(t, created)
	assert.ErrorContains(t, err, "internal: there was an error getting user by token")
}

func TestLoginWithErrorCreatingUser(t *testing.T) {
	oAuthRepository := new(oauthrepository.MockOAuthRepository)
	oAuthRepository.On("GenerateExternalToken").Return(
		"test_external_token",
		nil,
	)
	externalUserRepository := new(googleuserrepo.MockGoogleUserRepository)
	externalUserRepository.On("GetByToken").Return(
		testExternalUser,
		nil,
	)
	userUseCases := new(userusecases.MockUserUseCases)
	userUseCases.On("Create").Return(
		"",
		false,
		goerrors.New("there was an error creating user"),
	)

	oAuthUseCases := MakeOAuthUseCases(oAuthRepository, externalUserRepository, userUseCases)

	token, created, err := oAuthUseCases.Login("test_code")

	assert.Equal(t, "", token)
	assert.False(t, created)
	assert.ErrorContains(t, err, "internal: there was an error creating user")
}

func TestLoginWithSuccess(t *testing.T) {
	oAuthRepository := new(oauthrepository.MockOAuthRepository)
	oAuthRepository.On("GenerateExternalToken").Return(
		"test_external_token",
		nil,
	)
	externalUserRepository := new(googleuserrepo.MockGoogleUserRepository)
	externalUserRepository.On("GetByToken").Return(
		testExternalUser,
		nil,
	)
	userUseCases := new(userusecases.MockUserUseCases)
	userUseCases.On("Create").Return(
		"test_id",
		true,
		nil,
	)

	oAuthUseCases := MakeOAuthUseCases(oAuthRepository, externalUserRepository, userUseCases)

	token, created, err := oAuthUseCases.Login("test_code")

	assert.NotEmpty(t, token)
	assert.True(t, created)
	assert.Nil(t, err)
}
