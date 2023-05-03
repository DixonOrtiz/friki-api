package authusecases

import (
	authports "frikiapi/src/adapters/ports/use_cases/auth"
)

func MakeAuthUseCases(
	authRepository authports.AuthRepository,
	userUseCases authports.UserUseCases,
) AuthUseCasesInterface {
	return &AuthUseCases{
		AuthRepository: authRepository,
		UserUseCases:   userUseCases,
	}
}
