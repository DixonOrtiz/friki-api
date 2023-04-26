package types

import (
	authusecases "frikiapi/src/use_cases/auth"
	storeusecases "frikiapi/src/use_cases/store"
)

type UseCases struct {
	Auth  authusecases.AuthUseCasesInterface
	Store storeusecases.StoreUseCasesInterface
}
