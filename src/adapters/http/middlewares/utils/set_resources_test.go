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

func TestSetAddressResource(t *testing.T) {
	path := "/users/test_user_id/addresses/test_address_id"
	testAddressID := "test_address_id"

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.AddParam(permissions.ADDRESS_ID, testAddressID)

	resources := SetResources(path, c)

	assert.Len(t, resources, 1)
	assert.Equal(t, testAddressID, resources[permissions.ADDRESS])
}
