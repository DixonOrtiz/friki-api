package types

import (
	addressusecases "frikiapi/src/use_cases/addresses"
	oauthusecases "frikiapi/src/use_cases/oauth"
	permusecases "frikiapi/src/use_cases/permissions"
	userusecases "frikiapi/src/use_cases/users"
)

type UseCases struct {
	OAuth      oauthusecases.IOAuthUseCases
	User       userusecases.IUserUseCases
	Address    addressusecases.IAddressUseCases
	Permission permusecases.IPermissionUseCases
}
