package assembler

import (
	addresshttp "frikiapi/src/adapters/http/controllers/addresses"
	userhttp "frikiapi/src/adapters/http/controllers/users"
	mid "frikiapi/src/adapters/http/middlewares"
)

func (a *Assembler) setRoutes() {
	a.useMiddlewares()
	a.setOAuthRoutes()
	a.setUserRoutes()
	a.setAddressRoutes()
}

func (a *Assembler) useMiddlewares() {
	a.middlewares = mid.MakeMiddlewares()

	a.infraestructure.ProtectedRouter = a.infraestructure.Router.Group("/")
	a.infraestructure.ProtectedRouter.Use(a.middlewares.ValidateToken())
}

func (a *Assembler) setOAuthRoutes() {
	auth := a.infraestructure.Router.Group("/oauth")
	auth.GET("/login", a.controllers.OAuth.GoogleLogin)
	auth.GET("/callback", a.controllers.OAuth.GoogleCallback)
}

func (a *Assembler) setUserRoutes() {
	userhttp.SetRoutes(a.infraestructure.ProtectedRouter, a.controllers.User)
}

func (a *Assembler) setAddressRoutes() {
	addresshttp.SetRoutes(a.infraestructure.ProtectedRouter, a.controllers.Address)
}
