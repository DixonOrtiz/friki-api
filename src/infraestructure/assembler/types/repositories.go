package types

import (
	oauthrepo "frikiapi/src/adapters/repositories/oauth"
	storerepo "frikiapi/src/adapters/repositories/store"
	userrepo "frikiapi/src/adapters/repositories/user"
)

type Repositories struct {
	OAuth   oauthrepo.IOAuthRepository
	User    userrepo.IUserRepository
	Store   storerepo.IStoreRepository
}
