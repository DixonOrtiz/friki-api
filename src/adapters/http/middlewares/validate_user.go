package middlewares

import (
	"fmt"
	httputils "frikiapi/src/adapters/http/utils"
	"frikiapi/src/entities"
	httpinfra "frikiapi/src/infraestructure/http"
	"frikiapi/src/utils/consts"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m Middlewares) ValidateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		externalID := httputils.GetExternalID(c)

		user, err := m.UserUseCases.GetByExternalID(externalID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, httpinfra.Response{
				Error: fmt.Sprintf("%s: %s", consts.UNAUTHORIZED, err.Error()),
			})
			return
		}

		if user == (entities.User{}) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, httpinfra.Response{
				Error: fmt.Sprintf("%s: user with id: %d not found", consts.UNAUTHORIZED, user.ID),
			})
			return
		}

		c.Set("user_id", user.ID)
	}
}
