package utils

import (
	"frikiapi/src/utils/permissions"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSetEmptyResources(t *testing.T) {
	path := "/users/test_user_id"
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	resources := SetResources(path, c)

	assert.Len(t, resources, 0)
	assert.Empty(t, resources)
}

func TestSetEmptyResourceWithAddressesInPath(t *testing.T) {
	path := "/users/test_user_id/addresses"
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	resources := SetResources(path, c)

	assert.Len(t, resources, 0)
	assert.Empty(t, resources)
}

func TestSetEmptyResourceWithStoreInPath(t *testing.T) {
	path := "/users/test_user_id/stores"
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	resources := SetResources(path, c)

	assert.Len(t, resources, 0)
	assert.Empty(t, resources)
}

func TestSetAddressResource(t *testing.T) {
	path := "/users/test_user_id/addresses/test_address_id"
	testAddressID := "test_address_id"

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.AddParam(permissions.ADDRESS_ID, testAddressID)

	resources := SetResources(path, c)

	assert.Len(t, resources, 1)
	assert.Equal(t, testAddressID, resources[permissions.ADDRESS])
}

func TestSetStoreResource(t *testing.T) {
	path := "/users/test_user_id/stores/test_store_id"
	testStoreID := "test_store_id"

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.AddParam(permissions.STORE_ID, testStoreID)

	resources := SetResources(path, c)

	assert.Len(t, resources, 1)
	assert.Equal(t, testStoreID, resources[permissions.STORE])
}
