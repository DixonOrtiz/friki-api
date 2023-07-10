package permusecases

import (
	goerrors "errors"
	permrepo "frikiapi/src/adapters/repositories/permissions"
	"frikiapi/src/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePermissionWithErrorInCreate(t *testing.T) {
	permissionRepository := new(permrepo.MockPermissionRepository)
	permissionRepository.On("Create").Return(
		goerrors.New("there was an error creating the permission"),
	)
	permissionUseCases := MakePermissionUseCases(permissionRepository)

	permission, err := permissionUseCases.Create("test_user_id")

	assert.Empty(t, permission)
	assert.ErrorContains(t, err, "internal: there was an error creating the permission")
}

func TestCreatePermissionWithSuccess(t *testing.T) {
	permissionRepository := new(permrepo.MockPermissionRepository)
	permissionRepository.On("Create").Return(nil)
	permissionUseCases := MakePermissionUseCases(permissionRepository)

	userID := "test_user_id"

	permission, err := permissionUseCases.Create(userID)

	expectedPermission := entities.Permission{
		ID:        permission.ID,
		UserID:    userID,
		Addresses: []string{},
	}

	assert.Equal(t, expectedPermission, permission)
	assert.Nil(t, err)
}
