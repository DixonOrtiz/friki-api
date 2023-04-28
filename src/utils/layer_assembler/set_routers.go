package assembler

import mid "frikiapi/src/adapters/http/middlewares"

func (a *LayerAssembler) setRouters() {
	a.useMiddlewares()
	a.setAuthRouter()
	a.setStoreRouter()
}

func (a *LayerAssembler) useMiddlewares() {
	a.middlewares = mid.MakeMiddlewares(
		a.useCases.Store,
	)

	a.infraestructure.ProtectedRouter = a.infraestructure.Router.Group("/")
	a.infraestructure.ProtectedRouter.Use(a.middlewares.ValidateToken())
}

func (a *LayerAssembler) setAuthRouter() {
	auth := a.infraestructure.Router.Group("/auth")
	auth.GET("/login", a.controllers.Auth.GoogleLogin)
	auth.GET("/callback", a.controllers.Auth.GoogleCallback)
}

func (a *LayerAssembler) setStoreRouter() {
	store := a.infraestructure.ProtectedRouter.Group("/stores")
	protectedStore := store.Group("/:store_id")
	protectedStore.Use(a.middlewares.AuthorizeStore())

	store.POST("", a.controllers.Store.Create)
	protectedStore.POST("/addresses", a.controllers.Store.CreateAddress)
}
