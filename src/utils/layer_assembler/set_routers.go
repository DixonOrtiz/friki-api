package assembler

import mid "frikiapi/src/adapters/http/middlewares"

func (a *LayerAssembler) setRouters() {
	a.useMiddlewares()
	a.setAuthRouter()
	a.setUsersRouter()
}

func (a *LayerAssembler) useMiddlewares() {
	middlewares := mid.MakeMiddlewares(a.useCases.User)

	a.infraestructure.ProtectedRouter = a.infraestructure.Router.Group("/")
	a.infraestructure.ProtectedRouter.Use(middlewares.ValidateToken())
	// a.infraestructure.ProtectedRouter.Use(middlewares.ValidateUser())
}

func (a *LayerAssembler) setAuthRouter() {
	a.infraestructure.Router.GET("/login", a.controllers.Auth.GoogleLogin)
	a.infraestructure.Router.GET("/callback", a.controllers.Auth.GoogleCallback)
}

func (a *LayerAssembler) setUsersRouter() {
	a.infraestructure.ProtectedRouter.GET("/test-user", a.controllers.User.TestUser)
}
