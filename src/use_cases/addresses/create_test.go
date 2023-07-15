package addressusecases

import (
	goerrors "errors"
	addressrepo "frikiapi/src/adapters/repositories/addresses"
	permusecases "frikiapi/src/use_cases/permissions"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAddressWithErrorInCreate(t *testing.T) {
	addressRepository := new(addressrepo.MockAddressRepository)
	addressRepository.On("Create").Return(
		goerrors.New("there was an error creating the address"),
	)
	permissionUseCases := new(permusecases.MockPermissionUseCases)
	addressUseCases := MakeAddressUseCases(addressRepository, permissionUseCases)

	address, err := addressUseCases.Create(testAddress)

	assert.Empty(t, address)
	assert.ErrorContains(t, err, "internal: there was an error creating the address")
}

func TestCreateWithErrorCreatingPermission(t *testing.T) {
	addressRepository := new(addressrepo.MockAddressRepository)
	addressRepository.On("Create").Return(nil)
	permissionUseCases := new(permusecases.MockPermissionUseCases)
	permissionUseCases.On("AddResource").Return(
		goerrors.New("there was an error adding resource"),
	)
	addressUseCases := MakeAddressUseCases(addressRepository, permissionUseCases)

	address, err := addressUseCases.Create(testAddress)

	assert.Empty(t, address)
	assert.ErrorContains(t, err, "internal: there was an error adding resource")
}

func TestCreateAddressWithSuccess(t *testing.T) {
	addressRepository := new(addressrepo.MockAddressRepository)
	addressRepository.On("Create").Return(nil)
	permissionUseCases := new(permusecases.MockPermissionUseCases)
	permissionUseCases.On("AddResource").Return(nil)
	addressUseCases := MakeAddressUseCases(addressRepository, permissionUseCases)
	expectedAddress := testAddress

	address, err := addressUseCases.Create(testAddress)

	expectedAddress.ID = address.ID
	expectedAddress.CreatedAt = address.CreatedAt
	expectedAddress.UpdatedAt = address.UpdatedAt

	assert.Equal(t, expectedAddress, address)
	assert.Nil(t, err)
}
