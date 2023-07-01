package storehttp

import (
	"frikiapi/src/adapters/http/controllers/store/types"
	"frikiapi/src/adapters/http/utils"
	"frikiapi/src/entities"
	httpinfra "frikiapi/src/infraestructure/http"
	"frikiapi/src/utils/consts"
	"frikiapi/src/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (sc StoreControllers) CreateAddress(c *gin.Context) {
	storeID, err := utils.GetParam(c, "store_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, httpinfra.Response{
			Error: errors.New(consts.BAD_REQUEST, err),
		})
		return
	}

	var body types.CreateAddressDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, httpinfra.Response{
			Error: utils.FormatInputError(err),
		})
		return
	}

	addressID, err := sc.StoreUseCases.CreateAddress(entities.Address{
		StoreID:        storeID,
		Region:         body.Region,
		City:           body.City,
		Commune:        body.Commune,
		Street:         body.Street,
		Number:         body.Number,
		AdditionalInfo: body.AdditionalInfo,
	})
	if err != nil {
		c.JSON(errors.GetStatusByErr(err), httpinfra.Response{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, httpinfra.Response{
		Data: map[string]int{
			"address_id": addressID,
		},
	})
}
