package userusecases

import userports "frikiapi/src/adapters/ports/use_cases/user"

func MakeUserUseCases(
	userRepository userports.UserRepository,
) UserUseCasesInterface {
	return &UserUseCases{
		UserRepository: userRepository,
	}
}
