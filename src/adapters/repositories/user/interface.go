package userrepo

import "frikiapi/src/entities"

type UserRepositoryInterface interface {
	GetExternalUserByToken(token string) (entities.User, error)
}
