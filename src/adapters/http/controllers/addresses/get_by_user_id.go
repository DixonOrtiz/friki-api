package addresshttp

import (
	httpinfra "frikiapi/src/infraestructure/http"
	"frikiapi/src/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (co AddressControllers) GetByUserID(c *gin.Context) {
	userID := c.Param("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, httpinfra.Response{
			Error: errors.New(errors.BAD_REQUEST, "user_id is required in path").Error(),
		})
		return
	}

	addresses, err := co.AddressUseCases.GetByUserID(userID)
	if err != nil {
		c.JSON(errors.GetStatusByErr(err), httpinfra.Response{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, httpinfra.Response{
		Message: "addresses found successfully",
		Data:    addresses,
	})
}
