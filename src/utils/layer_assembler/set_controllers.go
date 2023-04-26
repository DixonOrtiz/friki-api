package assembler

import (
	authhttp "frikiapi/src/adapters/http/controllers/auth"
)

func (a *LayerAssembler) setControllers() {
	a.setAuthControllers()

}

func (a *LayerAssembler) setAuthControllers() {
	a.controllers.Auth = authhttp.MakeAuthControllers(
		a.useCases.Auth,
	)
}
