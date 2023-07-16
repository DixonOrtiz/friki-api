package utils

import (
	"frikiapi/src/utils/permissions"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetResources(c *gin.Context) map[string]string {
	resources := make(map[string]string)
	path := c.Request.URL.Path

	if strings.Contains(path, permissions.ADDRESSES) {
		resources[permissions.ADDRESSES] = c.Param("address_id")
	}

	return resources
}
