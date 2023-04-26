package types

import (
	authhttp "frikiapi/src/adapters/http/controllers/auth"
	storehttp "frikiapi/src/adapters/http/controllers/store"
)

type Controllers struct {
	Auth  authhttp.AuthControllers
	Store storehttp.StoreControllers
}
