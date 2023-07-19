package addressusecases

import (
	goerrors "errors"
	addressrepo "frikiapi/src/adapters/repositories/addresses"
	"frikiapi/src/entities"
	permusecases "frikiapi/src/use_cases/permissions"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetByUserIDWithError(t *testing.T) {
	addressRepository := new(addressrepo.MockAddressRepository)
	addressRepository.On("GetByUserID").Return(
		[]entities.Address{},
		goerrors.New("there was an error getting addresses for user"),
	)

	permissionUseCases := new(permusecases.MockPermissionUseCases)
	addressUseCases := MakeAddressUseCases(addressRepository, permissionUseCases)

	addresses, err := addressUseCases.GetByUserID("test_user_id")

	assert.Nil(t, addresses)
	assert.ErrorContains(t, err, "internal: there was an error getting addresses for user")
}

func TestGetByUserIDWithAddressNotFound(t *testing.T) {
	addressRepository := new(addressrepo.MockAddressRepository)
	addressRepository.On("GetByUserID").Return(
		[]entities.Address{},
		nil,
	)

	permissionUseCases := new(permusecases.MockPermissionUseCases)
	addressUseCases := MakeAddressUseCases(addressRepository, permissionUseCases)

	addresses, err := addressUseCases.GetByUserID("test_user_id")

	assert.Empty(t, addresses)
	assert.ErrorContains(t, err, "not_found: no address found with user_id test_user_id")
}

func TestGetByUserIDWithSuccess(t *testing.T) {
	testAddresses := []entities.Address{
		testAddress,
		testAddress,
	}

	addressRepository := new(addressrepo.MockAddressRepository)
	addressRepository.On("GetByUserID").Return(
		testAddresses,
		nil,
	)

	permissionUseCases := new(permusecases.MockPermissionUseCases)
	addressUseCases := MakeAddressUseCases(addressRepository, permissionUseCases)

	addresses, err := addressUseCases.GetByUserID("test_user_id")

	assert.Equal(t, testAddresses, addresses)
	assert.Nil(t, err)
}
