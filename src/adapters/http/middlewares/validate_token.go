package middlewares

import (
	"fmt"
	httpinfra "frikiapi/src/infraestructure/http"
	"frikiapi/src/utils/consts"
	jwtauth "frikiapi/src/utils/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (m Middlewares) ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header["Authorization"]

		if authorization == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, httpinfra.Response{
				Error: fmt.Sprintf("%s: token in authorization header is required", consts.UNAUTHORIZED),
			})
			return
		}

		tokenStr := strings.Split(authorization[0], " ")[1]

		token, err := jwtauth.ParseJWT(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, httpinfra.Response{
				Error: fmt.Sprintf("%s: %s", consts.UNAUTHORIZED, err.Error()),
			})
			return
		}

		externalID, err := jwtauth.GetExternalIDFromClaims(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, httpinfra.Response{
				Error: fmt.Sprintf("%s: %s", consts.UNAUTHORIZED, err.Error()),
			})
			return
		}

		c.Set("external_id", externalID)
		c.Next()
	}
}
