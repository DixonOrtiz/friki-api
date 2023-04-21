package userusecases

import (
	"errors"
	"fmt"
	userrepo "frikiapi/src/adapters/repositories/user"
	"frikiapi/src/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetByID(t *testing.T) {
	userRepository := new(userrepo.MockUserRepository)
	userRepository.On("GetByID").Return(testNormalUser, nil)
	userUseCases := MakeUserUseCases(userRepository)

	user, err := userUseCases.GetByID(testUserID)

	assert.Equal(t, user, testNormalUser)
	assert.Nil(t, err)
}

func TestErrorGettingUserByID(t *testing.T) {
	testError := "there was an error getting user by id"
	userRepository := new(userrepo.MockUserRepository)
	userRepository.On("GetByID").Return(entities.User{}, errors.New(testError))
	userUseCases := MakeUserUseCases(userRepository)

	user, err := userUseCases.GetByID(testUserID)

	assert.Equal(t, entities.User{}, user)
	assert.EqualError(t, err, fmt.Sprintf("internal: %s", testError))
}
