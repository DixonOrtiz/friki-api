package permusecases

import (
	goerrors "errors"
	permrepo "frikiapi/src/adapters/repositories/permissions"
	"frikiapi/src/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserIDAndJWTUserIDAreDifferent(t *testing.T) {
	permissionRepository := new(permrepo.MockPermissionRepository)
	permissionUseCases := MakePermissionUseCases(permissionRepository)

	err := permissionUseCases.Authorize("test_jwt_user_id", "test_user_id", make(map[string]string))

	assert.ErrorContains(t, err, "unauthorized: path user_id and token user_id are not the same")
}

func TestErrorGettingUser(t *testing.T) {
	permissionRepository := new(permrepo.MockPermissionRepository)
	permissionUseCases := MakePermissionUseCases(permissionRepository)
	permissionRepository.On("GetByUserID").Return(entities.Permission{}, "", goerrors.New("there was an error getting the user"))

	resources := map[string]string{"address": "test_address_id"}

	err := permissionUseCases.Authorize("test_user_id", "test_user_id", resources)

	assert.ErrorContains(t, err, "internal: there was an error getting the user")
}

func TestAddressAllowUserAndPass(t *testing.T) {
	permissionRepository := new(permrepo.MockPermissionRepository)
	permissionUseCases := MakePermissionUseCases(permissionRepository)
	permissionRepository.On("GetByUserID").Return(testPermission, "test_document", nil)

	resources := map[string]string{"address": "test_first_address_id"}

	err := permissionUseCases.Authorize("test_user_id", "test_user_id", resources)

	assert.Nil(t, err)
}

func TestAddressDoesNotAllowUser(t *testing.T) {
	permissionRepository := new(permrepo.MockPermissionRepository)
	permissionUseCases := MakePermissionUseCases(permissionRepository)
	permissionRepository.On("GetByUserID").Return(testPermission, "test_document", nil)

	resources := map[string]string{"address": "not_your_test_address_id"}

	err := permissionUseCases.Authorize("test_user_id", "test_user_id", resources)

	assert.ErrorContains(t, err, "unauthorized: the requested resource (address) is not owned by this user")
}

func TestResourcesAreEmptyAndPass(t *testing.T) {
	permissionRepository := new(permrepo.MockPermissionRepository)
	permissionUseCases := MakePermissionUseCases(permissionRepository)

	err := permissionUseCases.Authorize("test_user_id", "test_user_id", make(map[string]string))

	assert.Nil(t, err)
}
