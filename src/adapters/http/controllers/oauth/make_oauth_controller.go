package oauthhttp

import (
	ports "frikiapi/src/adapters/ports/http/controllers/oauth"
)

func MakeOAuthControllers(
	oAuthUseCases ports.OAuthUseCases,
) OAuthControllers {
	return OAuthControllers{
		OAuthUseCases: oAuthUseCases,
	}
}
