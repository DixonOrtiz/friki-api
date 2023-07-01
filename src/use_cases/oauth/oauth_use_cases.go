package authusecases

import (
	oauthrepository "frikiapi/src/adapters/repositories/oauth"
	userusecases "frikiapi/src/use_cases/user"
)


type OAuthUseCases struct {
	OAuthRepository oauthrepository.IOAuthRepository
	UserUseCases    userusecases.IUserUseCases
}
