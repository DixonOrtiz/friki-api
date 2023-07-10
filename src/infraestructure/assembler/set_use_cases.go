package assembler

import (
	oauthusecases "frikiapi/src/use_cases/oauth"
	userusecases "frikiapi/src/use_cases/users"
)

func (a *Assembler) setUseCases() {
	a.setUserUseCases()
	a.setAuthUseCases()

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
