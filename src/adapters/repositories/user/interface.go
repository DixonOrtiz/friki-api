package userrepo

import "frikiapi/src/entities"

type UserRepositoryInterface interface {
	GetExternalUserByToken(token string) (entities.User, error)
	Create(user entities.User) (int, error)
	GetByExternalID(externalID string) (entities.User, error)
}
