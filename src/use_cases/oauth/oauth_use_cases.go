package authusecases

import (
	oauthrepository "frikiapi/src/adapters/repositories/oauth"
	googleuserrepository "frikiapi/src/adapters/repositories/users/google"
	userusecases "frikiapi/src/use_cases/users"
)

type OAuthUseCases struct {
	OAuthRepository        oauthrepository.IOAuthRepository
	ExternalUserRepository googleuserrepository.IGoogleUserRepository
	UserUseCases           userusecases.IUserUseCases
}
