package userrepo

import "frikiapi/src/entities"

type UserRepositoryInterface interface {
	GetByExternalID(externalID string) (entities.User, error)
	Create(user entities.User) (int, error)
	Update(user entities.User) error
	GetExternalUserByToken(token string) (entities.User, error)
}
