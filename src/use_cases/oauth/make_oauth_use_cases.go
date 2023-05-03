package authusecases

import (
	oauthports "frikiapi/src/adapters/ports/use_cases/oauth"
)

func MakeOAuthUseCases(
	oAuthRepository oauthports.OAuthRepository,
	userUseCases oauthports.UserUseCases,
) AuthUseCasesInterface {
	return &OAuthUseCases{
		OAuthRepository: oAuthRepository,
		UserUseCases:    userUseCases,
	}
}
