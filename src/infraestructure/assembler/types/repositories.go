package types

import (
	addressrepo "frikiapi/src/adapters/repositories/addresses"
	oauthrepo "frikiapi/src/adapters/repositories/oauth"
	permrepo "frikiapi/src/adapters/repositories/permissions"
	storerepo "frikiapi/src/adapters/repositories/stores"
	userrepo "frikiapi/src/adapters/repositories/users"
	googleuserrepo "frikiapi/src/adapters/repositories/users/google"
)

type Repositories struct {
	OAuth        oauthrepo.IOAuthRepository
	User         userrepo.IUserRepository
	ExternalUser googleuserrepo.IGoogleUserRepository
	Address      addressrepo.IAddressRepository
	Permission   permrepo.IPermissionRepository
	Store 	     storerepo.IStoreRepository
}
