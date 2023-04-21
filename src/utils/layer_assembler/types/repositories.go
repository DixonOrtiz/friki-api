package types

import (
	authrepo "frikiapi/src/adapters/repositories/auth"
	userrepo "frikiapi/src/adapters/repositories/user"
)

type Repositories struct {
	Auth authrepo.AuthRepositoryInterface
	User userrepo.UserRepositoryInterface
}
