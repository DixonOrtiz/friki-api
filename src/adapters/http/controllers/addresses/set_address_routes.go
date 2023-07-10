package addresshttp

import "github.com/gin-gonic/gin"

func SetRoutes(router *gin.RouterGroup, controllers AddressControllers) {
	users := router.Group("/addresses")
	users.POST("", controllers.Create)
}
