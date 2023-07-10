package types

import (
	oauthrepo "frikiapi/src/adapters/repositories/oauth"
	userrepo "frikiapi/src/adapters/repositories/users"
	googleuserrepo "frikiapi/src/adapters/repositories/users/google"
)

type Repositories struct {
	OAuth        oauthrepo.IOAuthRepository
	User         userrepo.IUserRepository
	ExternalUser googleuserrepo.IGoogleUserRepository
}
