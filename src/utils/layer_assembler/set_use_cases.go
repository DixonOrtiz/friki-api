package assembler

import (
	authusecases "frikiapi/src/use_cases/auth"
	storeusecases "frikiapi/src/use_cases/store"
)

func (a *LayerAssembler) setUseCases() {
	a.setAuthUseCases()
	a.setStoreUseCases()
}

func (a *LayerAssembler) setAuthUseCases() {
	a.useCases.Auth = authusecases.MakeAuthUseCases(
		a.repositories.Auth,
		a.repositories.User,
	)
}

func (a *LayerAssembler) setStoreUseCases() {
	a.useCases.Store = storeusecases.MakeStoreUseCases(
		a.repositories.Store,
		a.repositories.Address,
	)
}
