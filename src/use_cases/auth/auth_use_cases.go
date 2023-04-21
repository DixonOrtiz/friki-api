package authusecases

import (
	authports "frikiapi/src/adapters/ports/use_cases/auth"
)

type AuthUseCases struct {
	AuthRepository authports.AuthRepository
	UserRepository authports.UserRepository
}
