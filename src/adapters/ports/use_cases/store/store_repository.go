package userports

import "frikiapi/src/entities"

type StoreRepository interface {
	Create(store entities.Store) (int, error)
	GetByExternalID(externalID string) (entities.Store, error)
	GetByID(ID int) (entities.Store, error)
}
