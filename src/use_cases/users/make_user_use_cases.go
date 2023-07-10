package userusecases

import userrepository "frikiapi/src/adapters/repositories/users"

func MakeUserUseCases(
	userRepository userrepository.IUserRepository,
) IUserUseCases {
	return &UserUseCases{
		UserRepository: userRepository,
	}
}
