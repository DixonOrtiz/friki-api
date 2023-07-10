package userhttp

import userusecases "frikiapi/src/use_cases/users"

func MakeUserControllers(userUseCases userusecases.IUserUseCases) UserControllers {
	return UserControllers{
		UserUseCases: userUseCases,
	}

}
