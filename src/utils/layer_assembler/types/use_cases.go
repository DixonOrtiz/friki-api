package types

import (
	oauthusecases "frikiapi/src/use_cases/oauth"
	storeusecases "frikiapi/src/use_cases/store"
	userusecases "frikiapi/src/use_cases/user"
)

type UseCases struct {
	OAuth oauthusecases.AuthUseCasesInterface
	Store storeusecases.StoreUseCasesInterface
	User  userusecases.UserUseCasesInterface
}
