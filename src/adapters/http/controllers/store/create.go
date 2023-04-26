package storehttp

import (
	"frikiapi/src/adapters/http/controllers/store/types"
	"frikiapi/src/entities"
	httpinfra "frikiapi/src/infraestructure/http"
	cutils "frikiapi/src/use_cases/store/utils/create"
	"frikiapi/src/utils"
	"frikiapi/src/utils/errors"
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

	storeID, err := sc.StoreUseCases.Create(cutils.CreateInput{
		Store: entities.Store{
			ExternalID:  body.ExternalID,
			Name:        body.Name,
			Description: body.Description,
		},
		Address: entities.Address{
			Region:         body.Address.Region,
			City:           body.Address.City,
			Commune:        body.Address.Commune,
			Street:         body.Address.Street,
			Number:         body.Address.Number,
			AdditionalInfo: body.Address.AdditionalInfo,
		},
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
