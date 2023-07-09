package userhttp

import "github.com/gin-gonic/gin"

func SetRoutes(router *gin.RouterGroup, controllers UserControllers) {
	users := router.Group("/users")
	users.PUT(":id", controllers.Update)
	users.GET(":id", controllers.GetByID)
}
