package assembler

import (
	oauthusecases "frikiapi/src/use_cases/oauth"
	storeusecases "frikiapi/src/use_cases/store"
	userusecases "frikiapi/src/use_cases/user"
)

func (a *LayerAssembler) setUseCases() {
	a.setUserUseCases()
	a.setAuthUseCases()
	a.setStoreUseCases()
}

func (a *LayerAssembler) setUserUseCases() {
	a.useCases.User = userusecases.MakeUserUseCases(
		a.repositories.User,
	)
}

func (a *LayerAssembler) setAuthUseCases() {
	a.useCases.OAuth = oauthusecases.MakeOAuthUseCases(
		a.repositories.OAuth,
		a.useCases.User,
	)
}

func (a *LayerAssembler) setStoreUseCases() {
	a.useCases.Store = storeusecases.MakeStoreUseCases(
		a.repositories.Store,
		a.repositories.Address,
	)
}
