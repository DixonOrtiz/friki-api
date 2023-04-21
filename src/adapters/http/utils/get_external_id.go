package httputils

import "github.com/gin-gonic/gin"

func GetExternalID(c *gin.Context) string {
	externalID, _ := c.Get("external_id")
	return externalID.(string)
}
