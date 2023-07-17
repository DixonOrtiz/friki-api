package input

import (
	"fmt"
	"frikiapi/src/utils/errors"

	"github.com/gin-gonic/gin"
)

func ValidateUpdateAddressParams(c *gin.Context) error {
	message := "%s is required in request param"

	userID := c.Param("user_id")
	if userID == "" {
		return errors.New(errors.BAD_REQUEST, fmt.Sprintf(message, "user_id"))
	}

	addressID := c.Param("address_id")
	if addressID == "" {
		return errors.New(errors.BAD_REQUEST, fmt.Sprintf(message, "address_id"))
	}

	return nil
}
