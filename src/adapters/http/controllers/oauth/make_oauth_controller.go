package oauthhttp

import (
	oauthusecases "frikiapi/src/use_cases/oauth"
)

func MakeOAuthControllers(
	oAuthUseCases oauthusecases.IOAuthUseCases,
) OAuthControllers {
	return OAuthControllers{OAuthUseCases: oAuthUseCases}
}
