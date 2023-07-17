package addresshttp

import (
	"frikiapi/src/adapters/http/controllers/addresses/input"
	httputils "frikiapi/src/adapters/http/utils"
	"frikiapi/src/entities"
	httpinfra "frikiapi/src/infraestructure/http"
	"frikiapi/src/utils/errors"
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

	err = input.ValidateUpdateAddressParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, httpinfra.Response{
			Error: err.Error,
		})
		return
	}

	err = co.AddressUseCases.Update(entities.Address{
		ID:               c.Param("address_id"),
		UserID:           c.Param("user_id"),
		Name:             addressToUpdate.Name,
		City:             addressToUpdate.City,
		Commune:          addressToUpdate.Commune,
		Street:           addressToUpdate.Street,
		Number:           addressToUpdate.Number,
		Apartment:        addressToUpdate.Apartment,
		Sector:           addressToUpdate.Sector,
		Type:             addressToUpdate.Type,
		ExtraInformation: addressToUpdate.ExtraInformation,
	})
	if err != nil {
		c.JSON(errors.GetStatusByErr(err), httpinfra.Response{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, httpinfra.Response{
		Message: "address updated successfully",
	})
}
