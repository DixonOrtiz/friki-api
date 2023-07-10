package userusecases

import userrepository "frikiapi/src/adapters/repositories/users"

type UserUseCases struct {
	UserRepository userrepository.IUserRepository
}
