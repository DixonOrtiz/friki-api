package types

import (
	addressrepo "frikiapi/src/adapters/repositories/address"
	authrepo "frikiapi/src/adapters/repositories/auth"
	storerepo "frikiapi/src/adapters/repositories/store"
	userrepo "frikiapi/src/adapters/repositories/user"
)

type Repositories struct {
	Auth    authrepo.AuthRepositoryInterface
	User    userrepo.UserRepositoryInterface
	Address addressrepo.AddressRepositoryInterface
	Store   storerepo.StoreRepositoryInterface
}
