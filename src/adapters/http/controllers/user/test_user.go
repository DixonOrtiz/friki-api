package userhttp

import (
	httpinfra "frikiapi/src/infraestructure/http"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (co *UserControllers) TestUser(c *gin.Context) {
	c.JSON(http.StatusOK, httpinfra.Response{
		Data: "OK",
	})
}
