package types

import (
	oauthrepo "frikiapi/src/adapters/repositories/oauth"
	userrepo "frikiapi/src/adapters/repositories/user"
	googleuserrepo "frikiapi/src/adapters/repositories/user/google"
)

type Repositories struct {
	OAuth        oauthrepo.IOAuthRepository
	User         userrepo.IUserRepository
	ExternalUser googleuserrepo.IGoogleUserRepository
}
