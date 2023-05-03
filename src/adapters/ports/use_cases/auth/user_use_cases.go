package authports

import "frikiapi/src/entities"

type UserUseCases interface {
	Create(user entities.User) (bool, error)
	GetExternalUserByToken(token string) (entities.User, error)
}
