package assembler

import (
	addresshttp "frikiapi/src/adapters/http/controllers/addresses"
	oauthhttp "frikiapi/src/adapters/http/controllers/oauth"
	storehttp "frikiapi/src/adapters/http/controllers/stores"
	userhttp "frikiapi/src/adapters/http/controllers/users"
)

func (a *Assembler) setControllers() {
	a.setAuthControllers()
	a.setUserControllers()
	a.setAddressControllers()
	a.setStoreControllers()
}

func (a *Assembler) setAuthControllers() {
	a.controllers.OAuth = oauthhttp.MakeOAuthControllers(
		a.useCases.OAuth,
	)
}

func (a *Assembler) setUserControllers() {
	a.controllers.User = userhttp.MakeUserControllers(
		a.useCases.User,
	)
}

func (a *Assembler) setAddressControllers() {
	a.controllers.Address = addresshttp.MakeAddressControllers(
		a.useCases.Address,
	)
}

func (a *Assembler) setStoreControllers() {
	a.controllers.Store = storehttp.MakeStoreControllers(
		a.useCases.Store,
	)
}
