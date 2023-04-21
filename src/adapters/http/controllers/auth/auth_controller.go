package authhttp

import authusecases "frikiapi/src/use_cases/auth"

type AuthControllers struct {
	AuthUseCases authusecases.AuthUseCasesInterface
}
