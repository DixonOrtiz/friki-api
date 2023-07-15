package addresshttp

import (
	httpinfra "frikiapi/src/infraestructure/http"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (co AddressControllers) GetByID(c *gin.Context) {
	tokenUserID, _ := c.Get("user_id")
	userID := c.Param("user_id")
	addressID := c.Param("address_id")

	resources := map[string]string{"address_id": addressID}

	err := co.AddressUseCases.GetByID(tokenUserID.(string), userID, resources)

	isAuthorized := "no está autorizado"

	if err == nil {
		isAuthorized = "está autorizado"
	}

	c.JSON(http.StatusOK, httpinfra.Response{
		Data: isAuthorized,
	})
}
