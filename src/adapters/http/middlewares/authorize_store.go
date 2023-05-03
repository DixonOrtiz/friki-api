package middlewares

import (
	"fmt"
	httputils "frikiapi/src/adapters/http/utils"
	httpinfra "frikiapi/src/infraestructure/http"
	"frikiapi/src/utils/consts"
	"frikiapi/src/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m Middlewares) AuthorizeStore() gin.HandlerFunc {
	return func(c *gin.Context) {
		externalID, _ := c.Get("external_id")

		storeID, err := httputils.GetParam(c, "store_id")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, httpinfra.Response{
				Error: fmt.Sprintf("%s: %e", consts.BAD_REQUEST, err),
			})
			return
		}

		c.Set("store_id", storeID)

		err = m.StoreUseCases.Authorize(
			externalID.(string),
			storeID,
		)
		if err != nil {
			c.AbortWithStatusJSON(
				errors.GetStatusByErr(err),
				httpinfra.Response{
					Error: err.Error(),
				},
			)
			return
		}

		c.Next()
	}
}
