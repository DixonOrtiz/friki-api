package addressusecases

import (
	goerrors "errors"
	addressrepo "frikiapi/src/adapters/repositories/addresses"
	"frikiapi/src/entities"
	permusecases "frikiapi/src/use_cases/permissions"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetByIDWithError(t *testing.T) {
	addressRepository := new(addressrepo.MockAddressRepository)
	addressRepository.On("GetByID").Return(
		testAddress,
		"test_address_doc",
		goerrors.New("there was an error getting the address"),
	)
	permissionUseCases := new(permusecases.MockPermissionUseCases)
	addressUseCases := MakeAddressUseCases(addressRepository, permissionUseCases)

	address, err := addressUseCases.GetByID(testAddress.ID)

	assert.Empty(t, address)
	assert.ErrorContains(t, err, "internal: there was an error getting the address")
}

func TestGetByIDWithAddressNotFound(t *testing.T) {
	addressRepository := new(addressrepo.MockAddressRepository)
	addressRepository.On("GetByID").Return(
		entities.Address{},
		"",
		nil,
	)
	permissionUseCases := new(permusecases.MockPermissionUseCases)
	addressUseCases := MakeAddressUseCases(addressRepository, permissionUseCases)

	address, err := addressUseCases.GetByID(testAddress.ID)

	assert.Empty(t, address)
	assert.ErrorContains(t, err, "not_found: address with id test_address_id not found")
}

func TestGetByIDWithSuccess(t *testing.T) {
	addressRepository := new(addressrepo.MockAddressRepository)
	addressRepository.On("GetByID").Return(
		testAddress,
		"test_document",
		nil,
	)
	permissionUseCases := new(permusecases.MockPermissionUseCases)
	addressUseCases := MakeAddressUseCases(addressRepository, permissionUseCases)

	address, err := addressUseCases.GetByID(testAddress.ID)

	assert.Equal(t, testAddress, address)
	assert.Nil(t, err)
}
