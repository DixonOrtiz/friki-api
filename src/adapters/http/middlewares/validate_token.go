package middlewares

import (
	"fmt"
	httpinfra "frikiapi/src/infraestructure/http"
	"frikiapi/src/utils/errors"
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
				Error: fmt.Sprintf("%s: token in authorization header is required", errors.UNAUTHORIZED),
			})
			return
		}

		tokenStr := strings.Split(authorization[0], " ")[1]

		token, err := jwtauth.ParseJWT(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, httpinfra.Response{
				Error: fmt.Sprintf("%s: %s", errors.UNAUTHORIZED, err.Error()),
			})
			return
		}

		userID, err := jwtauth.GetUserIDFromClaims(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, httpinfra.Response{
				Error: fmt.Sprintf("%s: %s", errors.UNAUTHORIZED, err.Error()),
			})
			return
		}

		if userID == "<nil>" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, httpinfra.Response{
				Error: fmt.Sprintf("%s: %s", errors.UNAUTHORIZED, "user_id is required in header token"),
			})
			return
		}

		c.Set("user_id", userID)
		c.Next()
	}
}
