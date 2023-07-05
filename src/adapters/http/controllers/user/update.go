package userhttp

import (
	"frikiapi/src/adapters/http/controllers/user/types"
	httputils "frikiapi/src/adapters/http/utils"
	"frikiapi/src/entities"
	httpinfra "frikiapi/src/infraestructure/http"
	"frikiapi/src/utils/consts"
	"frikiapi/src/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (co UserControllers) Update(c *gin.Context) {
	var body types.UpdateUserDTO

	userID := c.Param("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, httpinfra.Response{
			Error: errors.New(consts.BAD_REQUEST, errors.New(consts.BAD_REQUEST, "id is required")),
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

	err = co.UserUseCases.Update(userID, entities.User{
		Name:     body.Name,
		LastName: body.LastName,
		Email:    body.Email,
		Picture:  body.Picture,
	})
	if err != nil {
		c.JSON(errors.GetStatusByErr(err), httpinfra.Response{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, httpinfra.Response{
		Data: map[string]string{
			"user_id": userID,
		},
	})
}
