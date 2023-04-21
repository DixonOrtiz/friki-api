package authhttp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (co *AuthControllers) GoogleLogin(c *gin.Context) {
	url := co.AuthUseCases.GenerateLoginURL()
	c.Redirect(http.StatusSeeOther, url)
}
