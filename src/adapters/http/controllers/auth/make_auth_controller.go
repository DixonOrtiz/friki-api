package authhttp

import (
	ports "frikiapi/src/adapters/ports/http/controllers/auth"
)

func MakeAuthControllers(
	authUseCases ports.AuthUseCases,
) AuthControllers {
	return AuthControllers{
		AuthUseCases: authUseCases,
	}
}
