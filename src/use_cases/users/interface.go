package userusecases

import "frikiapi/src/entities"

type IUserUseCases interface {
	Create(user entities.User) (string, bool, error)
	Update(user entities.User) error
	GetByID(ID string) (entities.User, error)
}
