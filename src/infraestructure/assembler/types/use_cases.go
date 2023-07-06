package types

import (
	oauthusecases "frikiapi/src/use_cases/oauth"
	userusecases "frikiapi/src/use_cases/user"
)

type UseCases struct {
	OAuth oauthusecases.IOAuthUseCases
	User  userusecases.IUserUseCases
}