package userusecases

import (
	goerrors "errors"
	userrepo "frikiapi/src/adapters/repositories/users"
	"frikiapi/src/entities"
	permusecases "frikiapi/src/use_cases/permissions"
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
	permissionUseCases := new(permusecases.MockPermissionUseCases)
	userUseCases := MakeUserUseCases(userRepository, permissionUseCases)

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
	permissionUseCases := new(permusecases.MockPermissionUseCases)
	userUseCases := MakeUserUseCases(userRepository, permissionUseCases)

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
	permissionUseCases := new(permusecases.MockPermissionUseCases)
	userUseCases := MakeUserUseCases(userRepository, permissionUseCases)

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
	permissionUseCases := new(permusecases.MockPermissionUseCases)
	userUseCases := MakeUserUseCases(userRepository, permissionUseCases)

	err := userUseCases.Update(testUser)

	assert.Nil(t, err)
}
