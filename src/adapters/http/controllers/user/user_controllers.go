package userhttp

import (
	userports "frikiapi/src/adapters/ports/http/controllers/user"
)

type UserControllers struct {
	UserUseCases userports.UserUseCases
}
