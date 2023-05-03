package assembler

import (
	oauthhttp "frikiapi/src/adapters/http/controllers/oauth"
	storehttp "frikiapi/src/adapters/http/controllers/store"
)

func (a *LayerAssembler) setControllers() {
	a.setAuthControllers()
	a.setStoreControllers()

}

func (a *LayerAssembler) setAuthControllers() {
	a.controllers.OAuth = oauthhttp.MakeOAuthControllers(
		a.useCases.OAuth,
	)
}

func (a *LayerAssembler) setStoreControllers() {
	a.controllers.Store = storehttp.MakeStoreControllers(
		a.useCases.Store,
	)
}
