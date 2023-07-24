package storerepo

import "frikiapi/src/entities"

type IStoreRepository interface {
	Create(store entities.Store) error
	GetByID(ID string) (entities.Store, string, error)
}
