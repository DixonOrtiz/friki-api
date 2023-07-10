package assembler

import (
	addressusecases "frikiapi/src/use_cases/addresses"
	oauthusecases "frikiapi/src/use_cases/oauth"
	userusecases "frikiapi/src/use_cases/users"
)

func (a *Assembler) setUseCases() {
	a.setUserUseCases()
	a.setAuthUseCases()
	a.setAddressUseCases()
}

func (a *Assembler) setUserUseCases() {
	a.useCases.User = userusecases.MakeUserUseCases(
		a.repositories.User,
	)
}

func (a *Assembler) setAuthUseCases() {
	a.useCases.OAuth = oauthusecases.MakeOAuthUseCases(
		a.repositories.OAuth,
		a.repositories.ExternalUser,
		a.useCases.User,
	)
}

func (a *Assembler) setAddressUseCases() {
	a.useCases.Address = addressusecases.MakeAddressUseCases(
		a.repositories.Address,
	)
}
