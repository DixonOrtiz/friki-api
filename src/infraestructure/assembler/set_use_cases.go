package assembler

import (
	addressusecases "frikiapi/src/use_cases/addresses"
	oauthusecases "frikiapi/src/use_cases/oauth"
	permusecases "frikiapi/src/use_cases/permissions"
	storeusecases "frikiapi/src/use_cases/stores"
	userusecases "frikiapi/src/use_cases/users"
)

func (a *Assembler) setUseCases() {
	a.setPermissionUseCases()
	a.setUserUseCases()
	a.setAuthUseCases()
	a.setAddressUseCases()
	a.setStoreUseCases()
}

func (a *Assembler) setPermissionUseCases() {
	a.useCases.Permission = permusecases.MakePermissionUseCases(
		a.repositories.Permission,
	)
}

func (a *Assembler) setUserUseCases() {
	a.useCases.User = userusecases.MakeUserUseCases(
		a.repositories.User,
		a.useCases.Permission,
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
		a.useCases.Permission,
	)
}

func (a *Assembler) setStoreUseCases() {
	a.useCases.Store = storeusecases.MakeStoreUseCases(
		a.repositories.Store,
		a.useCases.Permission,
	)
}
