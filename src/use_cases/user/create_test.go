package userusecases

import (
	goerrors "errors"
	userrepo "frikiapi/src/adapters/repositories/user"
	"frikiapi/src/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateWithErrorInDoesExist(t *testing.T) {
	userRepository := new(userrepo.MockUserRepository)
	userRepository.On("GetByExternalID").Return(
		testUser,
		"test_user_doc",
		goerrors.New("there was an error verifing the user existence"),
	)
	userUseCases := MakeUserUseCases(userRepository)

	created, err := userUseCases.Create(testUser)

	assert.False(t, created)
	assert.ErrorContains(t, err, "internal: there was an error verifing the user existence")
}

func TestTryToCreateAnExistingUser(t *testing.T) {
	userRepository := new(userrepo.MockUserRepository)
	userRepository.On("GetByExternalID").Return(
		entities.User{ExternalID: "test_user_doc"},
		"",
		nil,
	)
	userUseCases := MakeUserUseCases(userRepository)

	created, err := userUseCases.Create(testUser)

	assert.False(t, created)
	assert.Nil(t, err)
}
