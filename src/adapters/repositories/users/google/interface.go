package googleuserrepo

import "frikiapi/src/entities"

type IGoogleUserRepository interface {
	GetByToken(token string) (entities.User, error)
}
