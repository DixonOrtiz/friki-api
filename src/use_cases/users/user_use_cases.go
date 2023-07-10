package userusecases

import (
	userrepository "frikiapi/src/adapters/repositories/users"
	permusecases "frikiapi/src/use_cases/permissions"
)

type UserUseCases struct {
	UserRepository     userrepository.IUserRepository
	PermissionUseCases permusecases.IPermissionUseCases
}
