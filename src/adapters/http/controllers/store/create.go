package storehttp

import (
	"frikiapi/src/adapters/http/controllers/store/types"
	httpinfra "frikiapi/src/infraestructure/http"
	"frikiapi/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (sc StoreControllers) Create(c *gin.Context) {
	var body types.CreateStoreDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, httpinfra.Response{
			Error: utils.FormatInputError(err),
		})
		return
	}

}
