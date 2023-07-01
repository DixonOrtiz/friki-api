package userusecases

import userrepository "frikiapi/src/adapters/repositories/user"


type UserUseCases struct {
	UserRepository userrepository.IUserRepository
}
