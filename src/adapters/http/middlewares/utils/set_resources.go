package utils

import (
	"frikiapi/src/utils/permissions"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetResources(path string, c *gin.Context) map[string]string {
	resources := make(map[string]string)

	if strings.Contains(path, permissions.ADDRESSES_PATH) {
		resources[permissions.ADDRESS] = c.Param(permissions.ADDRESS_ID)
	}

	return resources
}
