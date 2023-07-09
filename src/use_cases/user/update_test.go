package userusecases

import (
	goerrors "errors"
	userrepo "frikiapi/src/adapters/repositories/user"
	"frikiapi/src/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateUserWithErrorInDoesExist(t *testing.T) {
	userRepository := new(userrepo.MockUserRepository)
	userRepository.On("GetByID").Return(
		entities.User{},
		"",
		goerrors.New("there was an error verifing the user existence"),
	)
	userUseCases := MakeUserUseCases(userRepository)

	err := userUseCases.Update(testUser)

	assert.ErrorContains(t, err, "internal: there was an error verifing the user existence")
}

func TestUpdateAUserThatDoesNotExist(t *testing.T) {
	userRepository := new(userrepo.MockUserRepository)
	userRepository.On("GetByID").Return(
		entities.User{},
		"",
		nil,
	)
	userUseCases := MakeUserUseCases(userRepository)

	user := testUser
	user.ID = "test_id"

	err := userUseCases.Update(user)

	assert.ErrorContains(t, err, "not_found: user with id 'test_id' is not in the registers")
}

func TestUpdateUserWithErrorInUpdate(t *testing.T) {
	userRepository := new(userrepo.MockUserRepository)
	userRepository.On("GetByID").Return(
		testUser,
		document,
		nil,
	)
	userRepository.On("Update").Return(
		goerrors.New("there was an error updating user"),
	)
	userUseCases := MakeUserUseCases(userRepository)

	err := userUseCases.Update(testUser)

	assert.ErrorContains(t, err, "internal: there was an error updating user")
}

func TestUpdateUserWithSuccess(t *testing.T) {
	userRepository := new(userrepo.MockUserRepository)
	userRepository.On("GetByID").Return(
		testUser,
		document,
		nil,
	)
	userRepository.On("Update").Return(nil)
	userUseCases := MakeUserUseCases(userRepository)

	err := userUseCases.Update(testUser)

	assert.Nil(t, err)
}
