package addressusecases

import (
	goerrors "errors"
	addressrepo "frikiapi/src/adapters/repositories/addresses"
	"frikiapi/src/entities"
	permusecases "frikiapi/src/use_cases/permissions"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteAddressWithErrorInGetByID(t *testing.T) {
	addressRepository := new(addressrepo.MockAddressRepository)
	addressRepository.On("GetByID").Return(
		entities.Address{},
		"",
		goerrors.New("there was an error getting the address"),
	)
	permissionUseCases := new(permusecases.MockPermissionUseCases)
	addressUseCases := MakeAddressUseCases(addressRepository, permissionUseCases)

	err := addressUseCases.Delete(testAddressID)

	assert.EqualError(t, err, "internal: there was an error getting the address")
}

func TestDeleteAddressButAddressDoesNotExist(t *testing.T) {
	addressRepository := new(addressrepo.MockAddressRepository)
	addressRepository.On("GetByID").Return(
		entities.Address{},
		"",
		nil,
	)
	permissionUseCases := new(permusecases.MockPermissionUseCases)
	addressUseCases := MakeAddressUseCases(addressRepository, permissionUseCases)

	err := addressUseCases.Delete(testAddressID)

	assert.EqualError(t, err, "not_found: address with id test_address_id is not in the registers")
}

func TestDeleteAddressWithErrorInDelete(t *testing.T) {
	addressRepository := new(addressrepo.MockAddressRepository)
	addressRepository.On("GetByID").Return(
		entities.Address{
			ID: testAddressID,
		},
		"",
		nil,
	)
	addressRepository.On("Delete").Return(
		goerrors.New("there was an error deleting address"),
	)
	permissionUseCases := new(permusecases.MockPermissionUseCases)
	addressUseCases := MakeAddressUseCases(addressRepository, permissionUseCases)

	err := addressUseCases.Delete(testAddressID)

	assert.EqualError(t, err, "internal: there was an error deleting address")
}

func TestDeleteAddressWithErrorDeletingPermission(t *testing.T) {
	addressRepository := new(addressrepo.MockAddressRepository)
	addressRepository.On("GetByID").Return(
		entities.Address{
			ID: testAddressID,
		},
		"",
		nil,
	)
	addressRepository.On("Delete").Return(
		nil,
	)
	permissionUseCases := new(permusecases.MockPermissionUseCases)
	permissionUseCases.On("RemoveResource").Return(
		goerrors.New("there was an error removing resource"),
	)
	addressUseCases := MakeAddressUseCases(addressRepository, permissionUseCases)

	err := addressUseCases.Delete(testAddressID)

	assert.ErrorContains(t, err, "internal: there was an error removing resource")
}

func TestDeleteAddressWithSuccess(t *testing.T) {
	addressRepository := new(addressrepo.MockAddressRepository)
	addressRepository.On("GetByID").Return(
		entities.Address{
			ID: testAddressID,
		},
		"",
		nil,
	)
	addressRepository.On("Delete").Return(
		nil,
	)
	permissionUseCases := new(permusecases.MockPermissionUseCases)
	permissionUseCases.On("RemoveResource").Return(nil)
	addressUseCases := MakeAddressUseCases(addressRepository, permissionUseCases)

	err := addressUseCases.Delete(testAddressID)

	assert.Nil(t, err)
}
