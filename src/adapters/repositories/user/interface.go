package userrepo

import "frikiapi/src/entities"

type UserRepositoryInterface interface {
	GetByID(userID int) (entities.User, error)
	GetByExternalID(externalID string) (entities.User, error)
	GetExternalUserByToken(token string) (entities.User, error)
	Register(user entities.User) error
}
