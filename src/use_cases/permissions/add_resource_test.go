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
		entities.Permission{
			ID:        "test_id",
			UserID:    "user_id",
			Addresses: []string{},
		},
		"test_doc",
		nil,
	)
	permissionRepository.On("AddResource").Return(
		nil,
	)
	permissionUseCases := MakePermissionUseCases(permissionRepository)

	err := permissionUseCases.AddResource(permissions.ADDRESS, "test_user_id", "test_resource_id")

	assert.Nil(t, err)
}

func TestAddAddressResourceWithError(t *testing.T) {
	permissionRepository := new(permrepo.MockPermissionRepository)
	permissionRepository.On("GetByUserID").Return(
		entities.Permission{
			ID:        "test_id",
			UserID:    "user_id",
			Addresses: []string{},
		},
		"test_doc",
		nil,
	)
	permissionRepository.On("AddResource").Return(
		goerrors.New("there was an error adding permission resource"),
	)
	permissionUseCases := MakePermissionUseCases(permissionRepository)

	err := permissionUseCases.AddResource(permissions.ADDRESS, "test_user_id", "test_resource_id")

	assert.ErrorContains(t, err, "internal: there was an error adding permission resource")
}

func TestAddResourceWithErrorInResourceType(t *testing.T) {
	permissionRepository := new(permrepo.MockPermissionRepository)
	permissionRepository.On("GetByUserID").Return(
		entities.Permission{
			ID: "test_id",
		},
		"test_doc",
		nil,
	)
	permissionUseCases := MakePermissionUseCases(permissionRepository)

	err := permissionUseCases.AddResource("unsupported_resource", "test_user_id", "test_resource_id")

	assert.ErrorContains(t, err, "conflict: 'unsupported_resource' resource is not supported")
}
