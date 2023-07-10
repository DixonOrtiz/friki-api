package types

import (
	addressusecases "frikiapi/src/use_cases/addresses"
	oauthusecases "frikiapi/src/use_cases/oauth"
	userusecases "frikiapi/src/use_cases/users"
)

type UseCases struct {
	OAuth   oauthusecases.IOAuthUseCases
	User    userusecases.IUserUseCases
	Address addressusecases.IAddressUseCases
}
