package userusecases

import (
	goerrors "errors"
	userrepo "frikiapi/src/adapters/repositories/user"
	"frikiapi/src/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUserWithErrorInDoesExist(t *testing.T) {
	userRepository := new(userrepo.MockUserRepository)
	userRepository.On("GetByExternalID").Return(
		testUser,
		"test_user_doc",
		goerrors.New("there was an error verifing the user existence"),
	)
	userUseCases := MakeUserUseCases(userRepository)

	userID, created, err := userUseCases.Create(testUser)

	assert.Zero(t, userID)
	assert.False(t, created)
	assert.ErrorContains(t, err, "internal: there was an error verifing the user existence")
}

func TestTryToCreateAnExistingUser(t *testing.T) {
	userRepository := new(userrepo.MockUserRepository)
	userRepository.On("GetByExternalID").Return(
		entities.User{ID: "test_id"},
		"",
		nil,
	)
	userUseCases := MakeUserUseCases(userRepository)

	userID, created, err := userUseCases.Create(testUser)

	assert.Equal(t, "test_id", userID)
	assert.False(t, created)
	assert.Nil(t, err)
}

func TestCreateUserWithErrorInCreate(t *testing.T) {
	userRepository := new(userrepo.MockUserRepository)
	userRepository.On("GetByExternalID").Return(
		entities.User{},
		"",
		nil,
	)
	userRepository.On("Create").Return(goerrors.New("there was an error creating user"))
	userUseCases := MakeUserUseCases(userRepository)

	userID, created, err := userUseCases.Create(testUser)

	assert.Zero(t, userID)
	assert.False(t, created)
	assert.ErrorContains(t, err, "internal: there was an error creating user")
}

func TestCreateUserWithSuccess(t *testing.T) {
	userRepository := new(userrepo.MockUserRepository)
	userRepository.On("GetByExternalID").Return(
		entities.User{},
		"",
		nil,
	)
	userRepository.On("Create").Return(nil)
	userUseCases := MakeUserUseCases(userRepository)

	userID, created, err := userUseCases.Create(testUser)

	assert.NotZero(t, userID)
	assert.True(t, created)
	assert.Nil(t, err)
}
