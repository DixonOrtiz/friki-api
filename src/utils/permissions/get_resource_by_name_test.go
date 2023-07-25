package permissions

import (
	"frikiapi/src/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAddressesResourceByName(t *testing.T) {
	testAddresses := []string{"test_address_id_1", "test_address_id_2"}
	permission := entities.Permission{
		Addresses: testAddresses,
	}

	addresses := GetResourceByName(ADDRESS, permission)

	assert.Len(t, addresses, 2)
	assert.Equal(t, testAddresses, addresses)
}

func TestGetStoresResourceByName(t *testing.T) {
	testStores := []string{"test_first_store_id", "test_second_store_id", "test_third_store_id"}
	permission := entities.Permission{
		Stores: testStores,
	}

	stores := GetResourceByName(STORE, permission)

	assert.Len(t, stores, 3)
	assert.Equal(t, testStores, stores)
}

func TestGetResourceByNameAndGetNil(t *testing.T) {
	permission := entities.Permission{}
	resource := GetResourceByName("unsupported_resource", permission)
	assert.Nil(t, resource)
}
