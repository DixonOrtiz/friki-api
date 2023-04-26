package assembler

import mid "frikiapi/src/adapters/http/middlewares"

func (a *LayerAssembler) setRouters() {
	a.useMiddlewares()
	a.setAuthRouter()
}

func (a *LayerAssembler) useMiddlewares() {
	middlewares := mid.MakeMiddlewares()

	a.infraestructure.ProtectedRouter = a.infraestructure.Router.Group("/")
	a.infraestructure.ProtectedRouter.Use(middlewares.ValidateToken())
}

func (a *LayerAssembler) setAuthRouter() {
	auth := a.infraestructure.Router.Group("/auth")
	auth.GET("/login", a.controllers.Auth.GoogleLogin)
	auth.GET("/callback", a.controllers.Auth.GoogleCallback)
}
