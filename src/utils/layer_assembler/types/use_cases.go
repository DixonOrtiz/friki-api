package types

import (
	authusecases "frikiapi/src/use_cases/auth"
)

type UseCases struct {
	Auth authusecases.AuthUseCasesInterface
}
