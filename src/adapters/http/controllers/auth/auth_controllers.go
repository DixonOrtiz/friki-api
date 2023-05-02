package authhttp

import (
	contports "frikiapi/src/adapters/ports/http/controllers"
)

type AuthControllers struct {
	AuthUseCases contports.AuthUseCases
}
