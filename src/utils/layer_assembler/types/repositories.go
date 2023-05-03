package types

import (
	addressrepo "frikiapi/src/adapters/repositories/address"
	oauthrepo "frikiapi/src/adapters/repositories/oauth"
	storerepo "frikiapi/src/adapters/repositories/store"
	userrepo "frikiapi/src/adapters/repositories/user"
)

type Repositories struct {
	OAuth   oauthrepo.OAuthRepositoryInterface
	User    userrepo.UserRepositoryInterface
	Address addressrepo.AddressRepositoryInterface
	Store   storerepo.StoreRepositoryInterface
}
