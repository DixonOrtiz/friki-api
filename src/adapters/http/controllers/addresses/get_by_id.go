package addresshttp

import (
	httpinfra "frikiapi/src/infraestructure/http"
	"frikiapi/src/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (co AddressControllers) GetByID(c *gin.Context) {
	ID := c.Param("address_id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, httpinfra.Response{
			Error: errors.New(errors.BAD_REQUEST, "address_id is required in path").Error(),
		})
		return
	}

	address, err := co.AddressUseCases.GetByID(ID)
	if err != nil {
		c.JSON(errors.GetStatusByErr(err), httpinfra.Response{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, httpinfra.Response{
		Message: "address found successfully",
		Data:    address,
	})
}
