package storehttp

import (
	"frikiapi/src/adapters/http/controllers/stores/input"
	httputils "frikiapi/src/adapters/http/utils"
	"frikiapi/src/entities"
	httpinfra "frikiapi/src/infraestructure/http"
	"frikiapi/src/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (co StoreControllers) Create(c *gin.Context) {
	var body input.CreateStoreDTO
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

	_, err = co.StoreUseCases.Create(entities.Store{
		UserID:      userID.(string),
		Name:        body.Name,
		Description: body.Description,
	})
	if err != nil {
		c.JSON(errors.GetStatusByErr(err), httpinfra.Response{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, httpinfra.Response{
		Message: "store created successfully",
	})
}
