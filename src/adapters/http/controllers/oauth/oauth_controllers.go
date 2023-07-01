package oauthhttp

import (
	oauthusecases "frikiapi/src/use_cases/oauth"
)


type OAuthControllers struct {
	OAuthUseCases oauthusecases.IOAuthUseCases
}
