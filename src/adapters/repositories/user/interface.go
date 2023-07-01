package userrepo

import "frikiapi/src/entities"

type IUserRepository interface {
	GetByExternalID(externalID string) (entities.User, error)
	Create(user entities.User) (int, error)
	Update(user entities.User) error
}
