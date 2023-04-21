package types

import (
	authhttp "frikiapi/src/adapters/http/controllers/auth"
	userhttp "frikiapi/src/adapters/http/controllers/user"
)

type Controllers struct {
	Auth authhttp.AuthControllers
	User userhttp.UserControllers
}
