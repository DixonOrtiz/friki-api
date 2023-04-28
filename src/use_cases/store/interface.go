package storeusecases

import (
	"frikiapi/src/entities"
)

type StoreUseCasesInterface interface {
	Create(store entities.Store) (int, error)
	GetByExternalID(externalID string) (entities.Store, error)
	DoesExist(externalID string, ID int) (bool, error)
	CreateAddress(address entities.Address) (int, error)
}
