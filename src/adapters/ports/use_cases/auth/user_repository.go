package authports

import "frikiapi/src/entities"

type UserRepository interface {
	GetExternalUserByToken(token string) (entities.User, error)
}
