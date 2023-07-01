package authusecases

import (
	oauthrepository "frikiapi/src/adapters/repositories/oauth"
	userusecases "frikiapi/src/use_cases/user"
)

func MakeOAuthUseCases(
	oAuthRepository oauthrepository.IOAuthRepository,
	userUseCases userusecases.IUserUseCases,
) IOAuthUseCases {
	return &OAuthUseCases{
		OAuthRepository: oAuthRepository,
		UserUseCases:    userUseCases,
	}
}
