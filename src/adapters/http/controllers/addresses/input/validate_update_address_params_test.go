package input

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestValidateUpdateAddressParams(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Params = append(c.Params, gin.Param{
		Key:   "user_id",
		Value: "test_user_id",
	})
	c.Params = append(c.Params, gin.Param{
		Key:   "address_id",
		Value: "test_address_id",
	})

	err := ValidateUpdateAddressParams(c)

	assert.Nil(t, err)
}

func TestValidateUpdateAddressParamsUserIDIsRrequired(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Params = append(c.Params, gin.Param{
		Key:   "address_id",
		Value: "test_address_id",
	})

	err := ValidateUpdateAddressParams(c)

	assert.ErrorContains(t, err, "bad_request: user_id is required in request param")
}

func TestValidateUpdateAddressParamsAddressIDIsRrequired(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Params = append(c.Params, gin.Param{
		Key:   "user_id",
		Value: "test_user_id",
	})
	err := ValidateUpdateAddressParams(c)

	assert.ErrorContains(t, err, "bad_request: address_id is required in request param")
}
