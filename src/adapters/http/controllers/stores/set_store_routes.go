package storehttp

import "github.com/gin-gonic/gin"

func SetRoutes(router *gin.RouterGroup, controllers StoreControllers) {
	stores := router.Group("/stores")
	stores.POST("", controllers.Create)
}
