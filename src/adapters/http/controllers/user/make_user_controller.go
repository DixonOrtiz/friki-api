package userhttp

import (
	userports "frikiapi/src/adapters/ports/http/controllers/user"
)

func MakeUserControllers(userUseCases userports.UserUseCases) UserControllers {
	return UserControllers{
		UserUseCases: userUseCases,
	}

}
