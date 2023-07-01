package types

import (
	addressrepo "frikiapi/src/adapters/repositories/address"
	oauthrepo "frikiapi/src/adapters/repositories/oauth"
	storerepo "frikiapi/src/adapters/repositories/store"
	userrepo "frikiapi/src/adapters/repositories/user"
)

type Repositories struct {
	OAuth   oauthrepo.IOAuthRepository
	User    userrepo.IUserRepository
	Address addressrepo.IAddressRepository
	Store   storerepo.IStoreRepository
}
