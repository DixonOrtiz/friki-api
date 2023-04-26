package storerepo

import "frikiapi/src/entities"

type StoreRepositoryInterface interface {
	Create(store entities.Store) (int, error)
	GetByExternalID(externalID string) (entities.Store, error)
}
