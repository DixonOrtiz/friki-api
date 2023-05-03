package assembler

import (
	authusecases "frikiapi/src/use_cases/auth"
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
	a.useCases.Auth = authusecases.MakeAuthUseCases(
		a.repositories.Auth,
		a.useCases.User,
	)
}

func (a *LayerAssembler) setStoreUseCases() {
	a.useCases.Store = storeusecases.MakeStoreUseCases(
		a.repositories.Store,
		a.repositories.Address,
	)
}
