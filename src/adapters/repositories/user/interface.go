package userrepo

import "frikiapi/src/entities"

type IUserRepository interface {
	GetByExternalID(externalID string) (entities.User, string, error)
	Create(user entities.User) error
	Update(document string, user entities.User) error
}
