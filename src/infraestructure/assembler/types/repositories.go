package types

import (
	oauthrepo "frikiapi/src/adapters/repositories/oauth"
	storerepo "frikiapi/src/adapters/repositories/store"
	userrepo "frikiapi/src/adapters/repositories/user"
	googleuserrepo "frikiapi/src/adapters/repositories/user/google"
)

type Repositories struct {
	OAuth   oauthrepo.IOAuthRepository
	User    userrepo.IUserRepository
	Store   storerepo.IStoreRepository
	ExternalUser googleuserrepo.IGoogleUserRepository
}
