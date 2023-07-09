package userusecases

import "frikiapi/src/entities"

type IUserUseCases interface {
	Create(user entities.User) (bool, error)
	Update(user entities.User) error
}
