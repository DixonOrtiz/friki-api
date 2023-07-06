package userhttp

import "github.com/gin-gonic/gin"

func SetRoutes(router *gin.RouterGroup, controllers UserControllers) {
	users := router.Group("/users")
	users.PUT(":external_id", controllers.Update)
}
