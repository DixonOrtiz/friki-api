package permusecases

import (
	goerrors "errors"
	permrepo "frikiapi/src/adapters/repositories/permissions"
	"frikiapi/src/entities"
	"frikiapi/src/utils/permissions"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddResourceWithErrorGettingPermissionByID(t *testing.T) {
	permissionRepository := new(permrepo.MockPermissionRepository)
	permissionRepository.On("GetByUserID").Return(
		entities.Permission{},
		"",
		goerrors.New("there was an error getting the permission"),
	)
	permissionUseCases := MakePermissionUseCases(permissionRepository)

	err := permissionUseCases.AddResource(permissions.ADDRESS, "test_user_id", "test_resource_id")

	assert.ErrorContains(t, err, "internal: there was an error getting the permission")
}

func TestAddAddressResource(t *testing.T) {
	permissionRepository := new(permrepo.MockPermissionRepository)
	permissionRepository.On("GetByUserID").Return(
		emptyPermission,
		"test_doc",
		nil,
	)
	permissionRepository.On("UpdateResource").Return(
		nil,
	)
	permissionUseCases := MakePermissionUseCases(permissionRepository)

	err := permissionUseCases.AddResource(permissions.ADDRESS, "test_user_id", "test_resource_id")

	assert.Nil(t, err)
}

func TestAddAddressResourceWithError(t *testing.T) {
	permissionRepository := new(permrepo.MockPermissionRepository)
	permissionRepository.On("GetByUserID").Return(
		emptyPermission,
		"test_doc",
		nil,
	)
	permissionRepository.On("UpdateResource").Return(
		goerrors.New("there was an error adding permission resource"),
	)
	permissionUseCases := MakePermissionUseCases(permissionRepository)

	err := permissionUseCases.AddResource(permissions.ADDRESS, "test_user_id", "test_address_id")

	assert.ErrorContains(t, err, "internal: there was an error adding permission resource")
}

func TestAddStoreResource(t *testing.T) {
	permissionRepository := new(permrepo.MockPermissionRepository)
	permissionRepository.On("GetByUserID").Return(
		emptyPermission,
		"test_doc",
		nil,
	)
	permissionRepository.On("UpdateResource").Return(
		nil,
	)
	permissionUseCases := MakePermissionUseCases(permissionRepository)

	err := permissionUseCases.AddResource(permissions.ADDRESS, "test_user_id", "test_resource_id")

	assert.Nil(t, err)
}

func TestAddAddressResourceWithError(t *testing.T) {
	permissionRepository := new(permrepo.MockPermissionRepository)
	permissionRepository.On("GetByUserID").Return(
		emptyPermission,
		"test_doc",
		nil,
	)
	permissionRepository.On("UpdateResource").Return(
		goerrors.New("there was an error adding permission resource"),
	)
	permissionUseCases := MakePermissionUseCases(permissionRepository)

	err := permissionUseCases.AddResource(permissions.ADDRESS, "test_user_id", "test_address_id")

	assert.ErrorContains(t, err, "internal: there was an error adding permission resource")
}

func TestAddResourceWithErrorInResourceType(t *testing.T) {
	permissionRepository := new(permrepo.MockPermissionRepository)
	permissionRepository.On("GetByUserID").Return(
		testPermission,
		"test_doc",
		nil,
	)
	permissionUseCases := MakePermissionUseCases(permissionRepository)

	err := permissionUseCases.AddResource("unsupported_resource", "test_user_id", "test_resource_id")

	assert.ErrorContains(t, err, "conflict: 'unsupported_resource' resource is not supported")
}
