package assembler

import (
	authhttp "frikiapi/src/adapters/http/controllers/auth"
	userhttp "frikiapi/src/adapters/http/controllers/user"
)

func (a *LayerAssembler) setControllers() {
	a.setAuthControllers()
	a.setUserControllers()
}

func (a *LayerAssembler) setAuthControllers() {
	a.controllers.Auth = authhttp.MakeAuthControllers(
		a.useCases.Auth,
	)
}

func (a *LayerAssembler) setUserControllers() {
	a.controllers.User = userhttp.MakeUserControllers(
		a.useCases.User,
	)
}
