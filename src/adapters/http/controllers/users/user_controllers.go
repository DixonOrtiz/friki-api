package userhttp

import (
	userusecases "frikiapi/src/use_cases/users"
)

type UserControllers struct {
	UserUseCases userusecases.IUserUseCases
}
