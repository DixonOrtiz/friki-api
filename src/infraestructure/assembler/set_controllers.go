package assembler

import (
	addresshttp "frikiapi/src/adapters/http/controllers/addresses"
	oauthhttp "frikiapi/src/adapters/http/controllers/oauth"
	userhttp "frikiapi/src/adapters/http/controllers/users"
)

func (a *Assembler) setControllers() {
	a.setAuthControllers()
	a.setUserControllers()
	a.setAddressControllers()
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
