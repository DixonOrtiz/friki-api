package authports

import "frikiapi/src/entities"

type UserRepository interface {
	GetByExternalID(externalID string) (entities.User, error)
	GetExternalUserByToken(token string) (entities.User, error)
	Register(user entities.User) error
}
