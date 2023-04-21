package authhttp

import (
	httpinfra "frikiapi/src/infraestructure/http"
	"frikiapi/src/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CallbackResponseData struct {
	Token string `json:"token"`
}

func (co *AuthControllers) GoogleCallback(c *gin.Context) {
	code := c.Query("code")

	token, err := co.AuthUseCases.Login(code)
	if err != nil {
		c.JSON(errors.GetStatusByErr(err), httpinfra.Response{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, httpinfra.Response{
		Data: CallbackResponseData{
			Token: token,
		},
	})
}
