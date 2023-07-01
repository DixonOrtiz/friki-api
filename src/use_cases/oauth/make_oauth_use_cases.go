package authusecases

import (
	oauthrepository "frikiapi/src/adapters/repositories/oauth"
	googleuserrepository "frikiapi/src/adapters/repositories/user/google"
	userusecases "frikiapi/src/use_cases/user"
)

func MakeOAuthUseCases(
	oAuthRepository oauthrepository.IOAuthRepository,
	externalUserRepository googleuserrepository.IGoogleUserRepository,
	userUseCases userusecases.IUserUseCases,
) IOAuthUseCases {
	return &OAuthUseCases{
		OAuthRepository: oAuthRepository,
		ExternalUserRepository: externalUserRepository,
		UserUseCases:    userUseCases,
	}
}
