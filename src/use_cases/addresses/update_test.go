package addressusecases

import (
	"fmt"
	addressrepo "frikiapi/src/adapters/repositories/addresses"
	"frikiapi/src/entities"
	permusecases "frikiapi/src/use_cases/permissions"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateAddressWithErrorInGetByID(t *testing.T) {
	addressRepository := new(addressrepo.MockAddressRepository)
	addressRepository.On("GetByID").Return(
		entities.Address{},
		"",
		fmt.Errorf("there was an error getting the address"),
	)
	permissionUseCases := new(permusecases.MockPermissionUseCases)
	addressUseCases := MakeAddressUseCases(addressRepository, permissionUseCases)

	err := addressUseCases.Update(testAddress)

	assert.EqualError(t, err, "internal: there was an error getting the address")
}

func TestUpdateAddressWithNotFoundAddress(t *testing.T) {
	addressRepository := new(addressrepo.MockAddressRepository)
	addressRepository.On("GetByID").Return(
		entities.Address{},
		"",
		nil,
	)
	permissionUseCases := new(permusecases.MockPermissionUseCases)
	addressUseCases := MakeAddressUseCases(addressRepository, permissionUseCases)

	err := addressUseCases.Update(testAddress)

	assert.EqualError(t, err, "not_found: address with id 'test_address_id' is not in the registers")
}

func TestUpdateAddressWithErrorInUpdate(t *testing.T) {
	addressRepository := new(addressrepo.MockAddressRepository)
	addressRepository.On("GetByID").Return(
		testAddress,
		"test_document",
		nil,
	)
	addressRepository.On("Update").Return(
		fmt.Errorf("there was an error updating the address"),
	)
	permissionUseCases := new(permusecases.MockPermissionUseCases)
	addressUseCases := MakeAddressUseCases(addressRepository, permissionUseCases)

	err := addressUseCases.Update(testAddress)

	assert.EqualError(t, err, "internal: there was an error updating the address")
}

func TestUpdateAddressWithSuccess(t *testing.T) {
	addressRepository := new(addressrepo.MockAddressRepository)
	addressRepository.On("GetByID").Return(
		testAddress,
		"test_document",
		nil,
	)
	addressRepository.On("Update").Return(nil)
	permissionUseCases := new(permusecases.MockPermissionUseCases)
	addressUseCases := MakeAddressUseCases(addressRepository, permissionUseCases)

	err := addressUseCases.Update(testAddress)

	assert.Nil(t, err)
}
