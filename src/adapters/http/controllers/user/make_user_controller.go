package userhttp

import userusecases "frikiapi/src/use_cases/user"

func MakeUserControllers(userUseCases userusecases.UserUseCasesInterface) UserControllers {
	return UserControllers{
		UserUseCases: userUseCases,
	}
}
