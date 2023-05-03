package assembler

import (
	oauthhttp "frikiapi/src/adapters/http/controllers/oauth"
	storehttp "frikiapi/src/adapters/http/controllers/store"
	userhttp "frikiapi/src/adapters/http/controllers/user"
)

func (a *LayerAssembler) setControllers() {
	a.setAuthControllers()
	a.setUserControllers()
	a.setStoreControllers()

}

func (a *LayerAssembler) setAuthControllers() {
	a.controllers.OAuth = oauthhttp.MakeOAuthControllers(
		a.useCases.OAuth,
	)
}

func (a *LayerAssembler) setUserControllers() {
	a.controllers.User = userhttp.MakeUserControllers(
		a.useCases.User,
	)
}

func (a *LayerAssembler) setStoreControllers() {
	a.controllers.Store = storehttp.MakeStoreControllers(
		a.useCases.Store,
	)
}
