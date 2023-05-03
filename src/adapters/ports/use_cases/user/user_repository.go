package authports

import "frikiapi/src/entities"

type UserRepository interface {
	GetByExternalID(externalID string) (entities.User, error)
	Create(user entities.User) (int, error)
	Update(user entities.User) error
	GetExternalUserByToken(token string) (entities.User, error)
}
