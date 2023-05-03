package userports

import "frikiapi/src/entities"

type UserUseCases interface {
	Update(user entities.User) error
}
