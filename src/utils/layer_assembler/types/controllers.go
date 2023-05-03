package types

import (
	oauthttp "frikiapi/src/adapters/http/controllers/oauth"
	storehttp "frikiapi/src/adapters/http/controllers/store"
)

type Controllers struct {
	OAuth oauthttp.OAuthControllers
	Store storehttp.StoreControllers
}
