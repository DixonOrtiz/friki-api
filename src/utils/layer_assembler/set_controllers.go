package assembler

import (
	authhttp "frikiapi/src/adapters/http/controllers/auth"
	storehttp "frikiapi/src/adapters/http/controllers/store"
)

func (a *LayerAssembler) setControllers() {
	a.setAuthControllers()
	a.setStoreControllers()

}

func (a *LayerAssembler) setAuthControllers() {
	a.controllers.Auth = authhttp.MakeAuthControllers(
		a.useCases.Auth,
	)
}

func (a *LayerAssembler) setStoreControllers() {
	a.controllers.Store = storehttp.MakeStoreControllers(
		a.useCases.Store,
	)
}
