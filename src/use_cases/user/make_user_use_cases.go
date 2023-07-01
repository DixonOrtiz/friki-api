package userusecases

import userrepository "frikiapi/src/adapters/repositories/user"

func MakeUserUseCases(
	userRepository userrepository.IUserRepository,
) IUserUseCases {
	return &UserUseCases{
		UserRepository: userRepository,
	}
}
