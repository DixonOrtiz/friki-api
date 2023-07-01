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

	userID, err := httputils.GetParam(c, "user_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, httpinfra.Response{
			Error: errors.New(consts.BAD_REQUEST, err),
		})
		return
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, httpinfra.Response{
			Error: httputils.FormatInputError(err),
		})
		return
	}

	err = co.UserUseCases.Update(entities.User{
		ID:       userID,
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
		Data: map[string]int{
			"user_id": userID,
		},
	})
}
