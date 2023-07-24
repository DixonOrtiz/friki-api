package types

import (
	addressusecases "frikiapi/src/use_cases/addresses"
	oauthusecases "frikiapi/src/use_cases/oauth"
	permusecases "frikiapi/src/use_cases/permissions"
	storeusecases "frikiapi/src/use_cases/stores"
	userusecases "frikiapi/src/use_cases/users"
)

type UseCases struct {
	OAuth      oauthusecases.IOAuthUseCases
	User       userusecases.IUserUseCases
	Address    addressusecases.IAddressUseCases
	Permission permusecases.IPermissionUseCases
	Store      storeusecases.IStoreUseCases
}
