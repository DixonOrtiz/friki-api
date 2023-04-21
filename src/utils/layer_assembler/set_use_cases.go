package assembler

import (
	authusecases "frikiapi/src/use_cases/auth"
	userusecases "frikiapi/src/use_cases/user"
)

func (a *LayerAssembler) setUseCases() {
	a.setAuthUseCases()
	a.setUserUseCases()
}

func (a *LayerAssembler) setAuthUseCases() {
	a.useCases.Auth = authusecases.MakeAuthUseCases(
		a.repositories.Auth,
		a.repositories.User,
	)
}

func (a *LayerAssembler) setUserUseCases() {
	a.useCases.User = userusecases.MakeUserUseCases(
		a.repositories.User,
	)
}
