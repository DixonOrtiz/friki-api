package authusecases

import (
	authports "frikiapi/src/adapters/ports/use_cases/auth"
)

func MakeAuthUseCases(
	authRepository authports.AuthRepository,
	userRepository authports.UserRepository,
) AuthUseCasesInterface {
	return &AuthUseCases{
		AuthRepository: authRepository,
		UserRepository: userRepository,
	}
}
