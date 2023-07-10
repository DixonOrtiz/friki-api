package userusecases

import (
	userrepository "frikiapi/src/adapters/repositories/users"
	permusecases "frikiapi/src/use_cases/permissions"
)

func MakeUserUseCases(
	userRepository userrepository.IUserRepository,
	permissionUseCases permusecases.IPermissionUseCases,
) IUserUseCases {
	return &UserUseCases{
		UserRepository:     userRepository,
		PermissionUseCases: permissionUseCases,
	}
}
