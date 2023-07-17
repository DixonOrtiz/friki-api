package permusecases

import (
	goerrors "errors"
	permrepo "frikiapi/src/adapters/repositories/permissions"
	"frikiapi/src/entities"
	"frikiapi/src/utils/permissions"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveResourceWithErrorGettingPermissionByID(t *testing.T) {
	permissionRepository := new(permrepo.MockPermissionRepository)
	permissionRepository.On("GetByUserID").Return(
		entities.Permission{},
		"",
		goerrors.New("there was an error getting the permission"),
	)
	permissionUseCases := MakePermissionUseCases(permissionRepository)

	err := permissionUseCases.RemoveResource(permissions.ADDRESS, testUserID, testFirstAddressID)

	assert.ErrorContains(t, err, "internal: there was an error getting the permission")
}

func TestRemoveAddressResource(t *testing.T) {
	permissionRepository := new(permrepo.MockPermissionRepository)
	permissionRepository.On("GetByUserID").Return(
		testPermission,
		"test_doc",
		nil,
	)
	permissionRepository.On("UpdateResource").Return(
		nil,
	)
	permissionUseCases := MakePermissionUseCases(permissionRepository)

	err := permissionUseCases.RemoveResource(permissions.ADDRESS, testUserID, testFirstAddressID)

	assert.Nil(t, err)
}

func TestRemoveAddressResourceWithError(t *testing.T) {
	permissionRepository := new(permrepo.MockPermissionRepository)
	permissionRepository.On("GetByUserID").Return(
		testPermission,
		"test_doc",
		nil,
	)
	permissionRepository.On("UpdateResource").Return(
		goerrors.New("there was an error adding permission resource"),
	)
	permissionUseCases := MakePermissionUseCases(permissionRepository)

	err := permissionUseCases.RemoveResource(permissions.ADDRESS, testUserID, testFirstAddressID)

	assert.ErrorContains(t, err, "internal: there was an error adding permission resource")
}

func TestRemoveResourceWithErrorInResourceType(t *testing.T) {
	permissionRepository := new(permrepo.MockPermissionRepository)
	permissionRepository.On("GetByUserID").Return(
		testPermission,
		"test_doc",
		nil,
	)
	permissionUseCases := MakePermissionUseCases(permissionRepository)

	err := permissionUseCases.RemoveResource("unsupported_resource", testUserID, testFirstAddressID)

	assert.ErrorContains(t, err, "conflict: 'unsupported_resource' resource is not supported")
}
