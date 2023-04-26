package authrepo

import (
	authinfra "frikiapi/src/infraestructure/auth"
)

func MakeAuthRepository(config authinfra.OAuth2ConfigInterface) AuthRepositoryInterface {
	return AuthRepository{
		Config: config,
	}
}
