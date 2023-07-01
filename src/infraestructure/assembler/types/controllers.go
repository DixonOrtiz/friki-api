package types

import (
	oauthttp "frikiapi/src/adapters/http/controllers/oauth"
	storehttp "frikiapi/src/adapters/http/controllers/store"
	userhttp "frikiapi/src/adapters/http/controllers/user"
)

type Controllers struct {
	OAuth oauthttp.OAuthControllers
	Store storehttp.StoreControllers
	User  userhttp.UserControllers
}
