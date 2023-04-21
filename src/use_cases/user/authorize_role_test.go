package userusecases

import (
	userrepo "frikiapi/src/adapters/repositories/user"
	"frikiapi/src/utils/consts"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthorizeRole(t *testing.T) {
	userRepository := new(userrepo.MockUserRepository)
	userRepository.On("GetByID").Return(testSuperAdminUser, nil)
	userUseCases := MakeUserUseCases(userRepository)

	_, err := userUseCases.AuthorizeRole(testUserID, []int{consts.SUPER_ADMIN})

	assert.Nil(t, err)
}

func TestUnauthorizeRolePlural(t *testing.T) {
	userRepository := new(userrepo.MockUserRepository)
	userRepository.On("GetByID").Return(testNormalUser, nil)
	userUseCases := MakeUserUseCases(userRepository)

	_, err := userUseCases.AuthorizeRole(testUserID, []int{
		consts.SUPER_ADMIN,
		consts.ADMIN,
		consts.STAFF,
	})

	assert.EqualError(t, err, "unauthorized: this operation is allowed only for roles 1, 2 and 3, current role is 4")
}

func TestUnauthorizeRoleSingular(t *testing.T) {
	userRepository := new(userrepo.MockUserRepository)
	userRepository.On("GetByID").Return(testNormalUser, nil)
	userUseCases := MakeUserUseCases(userRepository)

	_, err := userUseCases.AuthorizeRole(testUserID, []int{
		consts.SUPER_ADMIN,
	})

	assert.EqualError(t, err, "unauthorized: this operation is allowed only for role 1, current role is 4")
}
