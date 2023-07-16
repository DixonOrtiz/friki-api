package middlewares

import (
	"frikiapi/src/adapters/http/middlewares/utils"
	httpinfra "frikiapi/src/infraestructure/http"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m Middlewares) Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenUserID, _ := c.Get("user_id")
		userID := c.Param("user_id")

		resources := utils.SetResources(c)

		err := m.PermissionUseCases.Authorize(
			tokenUserID.(string),
			userID,
			resources,
		)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, httpinfra.Response{
				Error: err.Error(),
			})
			return
		}

		c.Next()
	}
}
