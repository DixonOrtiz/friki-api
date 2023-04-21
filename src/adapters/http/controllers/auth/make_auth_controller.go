package authhttp

import authusecases "frikiapi/src/use_cases/auth"

func MakeAuthControllers(authUseCases authusecases.AuthUseCasesInterface) AuthControllers {
	return AuthControllers{
		AuthUseCases: authUseCases,
	}

}
