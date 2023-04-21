package types

import (
	authusecases "frikiapi/src/use_cases/auth"
	userusecases "frikiapi/src/use_cases/user"
)

type UseCases struct {
	Auth authusecases.AuthUseCasesInterface
	User userusecases.UserUseCasesInterface
}
