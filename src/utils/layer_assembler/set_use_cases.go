package assembler

import (
	authusecases "frikiapi/src/use_cases/auth"
)

func (a *LayerAssembler) setUseCases() {
	a.setAuthUseCases()
}

func (a *LayerAssembler) setAuthUseCases() {
	a.useCases.Auth = authusecases.MakeAuthUseCases(
		a.repositories.Auth,
		a.repositories.User,
	)
}
