package userusecases

import "frikiapi/src/entities"

type IUserUseCases interface {
	DoesExist(externalID string) (bool, error)
	Create(user entities.User) (bool, error)
	Update(ID string, user entities.User) error
}
