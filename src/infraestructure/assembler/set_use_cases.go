package assembler

import (
	oauthusecases "frikiapi/src/use_cases/oauth"
	storeusecases "frikiapi/src/use_cases/store"
	userusecases "frikiapi/src/use_cases/user"
)

func (a *Assembler) setUseCases() {
	a.setUserUseCases()
	a.setAuthUseCases()
	a.setStoreUseCases()
}

func (a *Assembler) setUserUseCases() {
	a.useCases.User = userusecases.MakeUserUseCases(
		a.repositories.User,
	)
}

func (a *Assembler) setAuthUseCases() {
	a.useCases.OAuth = oauthusecases.MakeOAuthUseCases(
		a.repositories.OAuth,
		a.useCases.User,
	)
}

func (a *Assembler) setStoreUseCases() {
	a.useCases.Store = storeusecases.MakeStoreUseCases(
		a.repositories.Store,
		a.repositories.Address,
	)
}
