package userhttp

import (
	"frikiapi/src/adapters/http/controllers/users/types"
	httputils "frikiapi/src/adapters/http/utils"
	"frikiapi/src/entities"
	httpinfra "frikiapi/src/infraestructure/http"
	"frikiapi/src/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (co UserControllers) Update(c *gin.Context) {
	var body types.UpdateUserDTO

	ID := c.Param("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, httpinfra.Response{
			Error: errors.New(errors.BAD_REQUEST, "id is required in path").Error(),
		})
		return
	}

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, httpinfra.Response{
			Error: httputils.FormatInputError(err),
		})
		return
	}

	err = co.UserUseCases.Update(entities.User{
		ID:       ID,
		Name:     body.Name,
		LastName: body.LastName,
		Picture:  body.Picture,
	})
	if err != nil {
		c.JSON(errors.GetStatusByErr(err), httpinfra.Response{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, httpinfra.Response{
		Message: "user updated successfully",
	})
}
