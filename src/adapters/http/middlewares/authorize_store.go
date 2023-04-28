package middlewares

import (
	"fmt"
	httputils "frikiapi/src/adapters/http/utils"
	httpinfra "frikiapi/src/infraestructure/http"
	"frikiapi/src/utils/consts"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m Middlewares) AuthorizeStore() gin.HandlerFunc {
	return func(c *gin.Context) {
		externalID, _ := c.Get("external_id")

		storeID, err := httputils.GetParam(c, "store_id", "int")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, httpinfra.Response{
				Error: fmt.Sprintf("%s: %e", consts.BAD_REQUEST, err),
			})
			return
		}

		store, err := m.StoreUseCases.GetByExternalID(externalID.(string))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, httpinfra.Response{
				Error: fmt.Sprintf("%s: %e", consts.INTERNAL, err),
			})
			return
		}

		if store.ID != storeID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, httpinfra.Response{
				Error: fmt.Sprintf("%s: %s", consts.UNAUTHORIZED, "only own stores can be modified"),
			})
			return

		}

		c.Next()
	}
}
