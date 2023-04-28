package httputils

import (
	"errors"
	"frikiapi/src/utils"

	"github.com/gin-gonic/gin"
)

func GetParam(c *gin.Context, param, parseTo string) (any, error) {
	paramStr := c.Param(param)

	switch parseTo {
	case "string":
		return paramStr, nil
	case "int":
		paramInt, err := utils.ParseStringToInt(paramStr)
		if err != nil {
			return nil, err
		}
		return paramInt, nil

	default:
		return nil, errors.New("param should be string or int")
	}
}
