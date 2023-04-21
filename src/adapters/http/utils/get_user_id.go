package httputils

import "github.com/gin-gonic/gin"

func GetUserID(c *gin.Context) int {
	userID, _ := c.Get("user_id")
	return userID.(int)
}
