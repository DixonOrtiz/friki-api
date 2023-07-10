package userusecases

import (
	goerrors "errors"
	userrepo "frikiapi/src/adapters/repositories/users"
	"frikiapi/src/entities"
	permusecases "frikiapi/src/use_cases/permissions"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindByIDWithError(t *testing.T) {
	userRepository := new(userrepo.MockUserRepository)
	userRepository.On("GetByID").Return(
		testUser,
		"test_user_doc",
		goerrors.New("there was an error getting the user"),
	)
	permissionUseCases := new(permusecases.MockPermissionUseCases)
	userUseCases := MakeUserUseCases(userRepository, permissionUseCases)

	user, err := userUseCases.GetByID(testUser.ID)

	assert.Empty(t, user)
	assert.ErrorContains(t, err, "internal: there was an error getting the user")
}

func TestFindByIDWithUserNotFound(t *testing.T) {
	userRepository := new(userrepo.MockUserRepository)
	userRepository.On("GetByID").Return(
		entities.User{},
		"",
		nil,
	)
	permissionUseCases := new(permusecases.MockPermissionUseCases)
	userUseCases := MakeUserUseCases(userRepository, permissionUseCases)

	user, err := userUseCases.GetByID(testUser.ID)

	assert.Empty(t, user)
	assert.ErrorContains(t, err, "not_found: user with id 'test_id' does not exist")
}

func TestFindByIDWithSuccess(t *testing.T) {
	userRepository := new(userrepo.MockUserRepository)
	userRepository.On("GetByID").Return(
		testUser,
		"",
		nil,
	)
	permissionUseCases := new(permusecases.MockPermissionUseCases)
	userUseCases := MakeUserUseCases(userRepository, permissionUseCases)

	user, err := userUseCases.GetByID(testUser.ID)

	assert.Equal(t, testUser, user)
	assert.Nil(t, err)
}
