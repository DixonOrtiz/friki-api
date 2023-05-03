package userusecases

import "frikiapi/src/entities"

type UserUseCasesInterface interface {
	DoesExist(externalID string) (bool, error)
	Create(user entities.User) (bool, error)
	Update(user entities.User) error
	GetExternalUserByToken(token string) (entities.User, error)
}
