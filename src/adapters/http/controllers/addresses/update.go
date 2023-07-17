package addresshttp

import (
	"frikiapi/src/adapters/http/controllers/addresses/input"
	httputils "frikiapi/src/adapters/http/utils"
	httpinfra "frikiapi/src/infraestructure/http"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (co AddressControllers) Update(c *gin.Context) {
	var addressToUpdate input.UpdateAddressDTO

	err := c.ShouldBindJSON(&addressToUpdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, httpinfra.Response{
			Error: httputils.FormatInputError(err),
		})
		return
	}

	// _, err = co.AddressUseCases.Update(addressToUpdate)
	// if err != nil {
	// 	c.JSON(errors.GetStatusByErr(err), httpinfra.Response{
	// 		Error: err.Error(),
	// 	})
	// 	return
	// }

	c.JSON(http.StatusOK, httpinfra.Response{
		Message: "address updated successfully",
	})
}
