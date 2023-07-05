package userrepo

import "frikiapi/src/entities"

type IUserRepository interface {
	GetByExternalID(externalID string) (entities.User, error)
	Create(user entities.User) error
	Update(ID string, user entities.User) error
}
