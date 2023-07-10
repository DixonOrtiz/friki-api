package addressusecases

import (
	goerrors "errors"
	addressrepo "frikiapi/src/adapters/repositories/addresses"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAddressWithErrorInCreate(t *testing.T) {
	addressRepository := new(addressrepo.MockAddressRepository)
	addressRepository.On("Create").Return(
		goerrors.New("there was an error creating the address"),
	)
	addressUseCases := MakeAddressUseCases(addressRepository)

	address, err := addressUseCases.Create(testAddress)

	assert.Empty(t, address)
	assert.ErrorContains(t, err, "internal: there was an error creating the address")
}

func TestCreateAddressWithSuccess(t *testing.T) {
	addressRepository := new(addressrepo.MockAddressRepository)
	addressRepository.On("Create").Return(nil)
	addressUseCases := MakeAddressUseCases(addressRepository)
	expectedAddress := testAddress

	address, err := addressUseCases.Create(testAddress)

	expectedAddress.ID = address.ID
	expectedAddress.CreatedAt = address.CreatedAt
	expectedAddress.UpdatedAt = address.UpdatedAt

	assert.Equal(t, expectedAddress, address)
	assert.Nil(t, err)
}
