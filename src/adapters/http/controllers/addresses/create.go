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

func (co AddressControllers) Create(c *gin.Context) {
	var body input.CreateAddressDTO
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusBadRequest, httpinfra.Response{
			Error: errors.New(
				errors.BAD_REQUEST,
				"user_id is required in header auth token",
			)})
		return
	}

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, httpinfra.Response{
			Error: httputils.FormatInputError(err),
		})
		return
	}

	_, err = co.AddressUseCases.Create(entities.Address{
		UserID:           userID.(string),
		Name:             body.Name,
		City:             body.City,
		Commune:          body.Commune,
		Street:           body.Street,
		Number:           body.Number,
		Apartment:        body.Apartment,
		Sector:           body.Sector,
		Type:             body.Type,
		ExtraInformation: body.ExtraInformation,
	})
	if err != nil {
		c.JSON(errors.GetStatusByErr(err), httpinfra.Response{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, httpinfra.Response{
		Message: "address created successfully",
	})
}
