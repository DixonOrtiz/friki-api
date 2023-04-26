package types

import (
	authhttp "frikiapi/src/adapters/http/controllers/auth"
)

type Controllers struct {
	Auth authhttp.AuthControllers
}
