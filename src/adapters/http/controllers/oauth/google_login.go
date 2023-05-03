package oauthhttp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (co *OAuthControllers) GoogleLogin(c *gin.Context) {
	url := co.OAuthUseCases.GenerateLoginURL()
	c.Redirect(http.StatusSeeOther, url)
}
