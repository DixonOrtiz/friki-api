package userhttp

import "github.com/gin-gonic/gin"

func SetRoutes(router *gin.RouterGroup, controllers UserControllers) {
	users := router.Group("/users/:user_id")
	users.PUT("", controllers.Update)
	users.GET("", controllers.GetByID)
}
