package addresshttp

import (
	httpinfra "frikiapi/src/infraestructure/http"
	"frikiapi/src/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (co AddressControllers) GetByID(c *gin.Context) {
	// tokenUserID, _ := c.Get("user_id")
	// userID := c.Param("user_id")
	// addressID := c.Param("address_id")

	// resources := map[string]string{"address_id": addressID}

	// err := co.AddressUseCases.GetByID(tokenUserID.(string), userID, resources)

	// isAuthorized := "no está autorizado"

	// if err == nil {
	// 	isAuthorized = "está autorizado"
	// }

	// c.JSON(http.StatusOK, httpinfra.Response{
	// 	Data: isAuthorized,
	// })
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
