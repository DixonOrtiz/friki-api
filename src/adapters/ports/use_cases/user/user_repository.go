package userports

import "frikiapi/src/entities"

type UserRepository interface {
	GetByID(userID int) (entities.User, error)
	GetByExternalID(externalID string) (entities.User, error)
}
