package types

import (
	authusecases "frikiapi/src/use_cases/auth"
	storeusecases "frikiapi/src/use_cases/store"
	userusecases "frikiapi/src/use_cases/user"
)

type UseCases struct {
	Auth  authusecases.AuthUseCasesInterface
	Store storeusecases.StoreUseCasesInterface
	User  userusecases.UserUseCasesInterface
}
