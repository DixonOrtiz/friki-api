package storeusecases

import (
	"frikiapi/src/entities"
)

type IStoreUseCases interface {
	Authorize(externalID string, storeID int) error
	Create(store entities.Store) (int, error)
	Update(store entities.Store) error
	GetByExternalID(externalID string) (entities.Store, error)
	DoesExist(externalID string, ID int) (bool, error)
	CreateAddress(address entities.Address) (int, error)
}
