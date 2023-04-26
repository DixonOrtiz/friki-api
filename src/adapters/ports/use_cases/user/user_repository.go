package userports

import "frikiapi/src/entities"

type UserRepository interface {
	GetExternalUserByToken(externalID string) (entities.User, error)
}
