package types

import (
	oauthttp "frikiapi/src/adapters/http/controllers/oauth"
	userhttp "frikiapi/src/adapters/http/controllers/users"
)

type Controllers struct {
	OAuth oauthttp.OAuthControllers
	User  userhttp.UserControllers
}
