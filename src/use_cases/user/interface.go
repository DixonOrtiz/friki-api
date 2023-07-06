package userusecases

import "frikiapi/src/entities"

type IUserUseCases interface {
	DoesExist(externalID string) (bool, string, error)
	Create(user entities.User) (bool, error)
	Update(user entities.User) error
}
