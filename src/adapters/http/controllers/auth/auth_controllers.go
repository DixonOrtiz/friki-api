package authhttp

import (
	ports "frikiapi/src/adapters/ports/http/controllers/auth"
)

type AuthControllers struct {
	AuthUseCases ports.AuthUseCases
}
