package authusecases

import (
	oauthports "frikiapi/src/adapters/ports/use_cases/oauth"
)

type OAuthUseCases struct {
	OAuthRepository oauthports.OAuthRepository
	UserUseCases    oauthports.UserUseCases
}
