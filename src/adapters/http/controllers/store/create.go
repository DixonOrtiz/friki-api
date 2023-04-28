package storehttp

import (
	"frikiapi/src/adapters/http/controllers/store/types"
	"frikiapi/src/entities"
	httpinfra "frikiapi/src/infraestructure/http"
	"frikiapi/src/utils"
	"frikiapi/src/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (sc StoreControllers) Create(c *gin.Context) {
	var body types.CreateStoreDTO
	externalID, _ := c.Get("external_id")

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, httpinfra.Response{
			Error: utils.FormatInputError(err),
		})
		return
	}

	storeID, err := sc.StoreUseCases.Create(entities.Store{
		ExternalID:  externalID.(string),
		Name:        body.Name,
		Description: body.Description,
	})
	if err != nil {
		c.JSON(errors.GetStatusByErr(err), httpinfra.Response{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, httpinfra.Response{
		Data: map[string]int{
			"store_id": storeID,
		},
	})
}
