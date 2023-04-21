package userusecases

import (
	"frikiapi/src/entities"
)

type UserUseCasesInterface interface {
	GetByID(userID int) (entities.User, error)
	GetByExternalID(externalID string) (entities.User, error)
	AuthorizeRole(userID int, roles []int) (entities.User, error)
}
