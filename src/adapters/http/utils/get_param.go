package utils

import (
	"frikiapi/src/utils/types"

	"github.com/gin-gonic/gin"
)

func GetParam(c *gin.Context, param string) (int, error) {
	paramStr := c.Param(param)

	paramInt, err := types.ParseStringToInt(paramStr)
	if err != nil {
		return 0, err
	}
	return paramInt, nil

}
