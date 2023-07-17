package addresshttp

import (
	"frikiapi/src/adapters/http/controllers/addresses/input"
	httpinfra "frikiapi/src/infraestructure/http"
	"frikiapi/src/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (co AddressControllers) Delete(c *gin.Context) {
	err := input.ValidateUserIDAndAddressIDParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, httpinfra.Response{
			Error: err.Error,
		})
		return
	}

	err = co.AddressUseCases.Delete(c.Param("address_id"))
	if err != nil {
		c.JSON(errors.GetStatusByErr(err), httpinfra.Response{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, httpinfra.Response{
		Message: "address deleted successfully",
	})
}
