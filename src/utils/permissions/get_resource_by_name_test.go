package permissions

import (
	"frikiapi/src/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetResourceByName(t *testing.T) {
	testAddresses := []string{"test_address_id_1", "test_address_id_2"}
	permission := entities.Permission{
		Addresses: testAddresses,
	}

	addresses := GetResourceByName(ADDRESS, permission)

	assert.Len(t, addresses, 2)
	assert.Equal(t, testAddresses, addresses)
}

func TestGetResourceByNameAndGetNil(t *testing.T) {
	permission := entities.Permission{}
	resource := GetResourceByName("unsupported_resource", permission)
	assert.Nil(t, resource)
}
