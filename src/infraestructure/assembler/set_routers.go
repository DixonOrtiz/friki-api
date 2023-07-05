package assembler

import mid "frikiapi/src/adapters/http/middlewares"

func (a *Assembler) setRouters() {
	a.useMiddlewares()
	a.setOAuthRouter()
	a.setUserRouter()
}

func (a *Assembler) useMiddlewares() {
	a.middlewares = mid.MakeMiddlewares()

	a.infraestructure.ProtectedRouter = a.infraestructure.Router.Group("/")
	a.infraestructure.ProtectedRouter.Use(a.middlewares.ValidateToken())
}

func (a *Assembler) setOAuthRouter() {
	auth := a.infraestructure.Router.Group("/oauth")
	auth.GET("/login", a.controllers.OAuth.GoogleLogin)
	auth.GET("/callback", a.controllers.OAuth.GoogleCallback)
}

func (a *Assembler) setUserRouter() {
	users := a.infraestructure.ProtectedRouter.Group("/users")
	users.PUT(":user_id", a.controllers.User.Update)
}
