package addresshttp

import "github.com/gin-gonic/gin"

func SetRoutes(router *gin.RouterGroup, controllers AddressControllers) {
	addresses := router.Group("/users/:user_id/addresses")
	addresses.POST("", controllers.Create)
}
