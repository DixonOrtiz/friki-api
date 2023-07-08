package userusecases

import (
	goerrors "errors"
	userrepo "frikiapi/src/adapters/repositories/user"
	"frikiapi/src/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoesExist(t *testing.T) {
	userRepository := new(userrepo.MockUserRepository)
	userRepository.On("GetByExternalID").Return(testUser, "test_user_doc", nil)

	userUseCases := MakeUserUseCases(userRepository)

	exist, document, err := userUseCases.DoesExist("test_external_id")

	assert.True(t, exist)
	assert.Equal(t, "test_user_doc", document)
	assert.Nil(t, err)
}

func TestDoesExistWithError(t *testing.T) {
	userRepository := new(userrepo.MockUserRepository)
	userRepository.On("GetByExternalID").Return(
		testUser,
		"test_user_doc",
		goerrors.New("there was an error interacting with db"),
	)

	userUseCases := MakeUserUseCases(userRepository)

	exist, document, err := userUseCases.DoesExist("test_external_id")

	assert.False(t, exist)
	assert.Zero(t, document)
	assert.ErrorContains(t, err, "internal: there was an error interacting with db")
}

func TestDoesNotExist(t *testing.T) {
	userRepository := new(userrepo.MockUserRepository)
	userRepository.On("GetByExternalID").Return(
		entities.User{},
		"",
		nil,
	)

	userUseCases := MakeUserUseCases(userRepository)

	exist, document, err := userUseCases.DoesExist("test_external_id")

	assert.False(t, exist)
	assert.Zero(t, document)
	assert.Nil(t, err)
}
