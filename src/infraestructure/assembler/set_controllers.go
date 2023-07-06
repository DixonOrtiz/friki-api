package assembler

import (
	oauthhttp "frikiapi/src/adapters/http/controllers/oauth"
	userhttp "frikiapi/src/adapters/http/controllers/users"
)

func (a *Assembler) setControllers() {
	a.setAuthControllers()
	a.setUserControllers()
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
