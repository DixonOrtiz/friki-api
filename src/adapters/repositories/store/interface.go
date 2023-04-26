package storerepo

import "frikiapi/src/entities"

type StoreRepositoryInterface interface {
	Create(store entities.Store, addressID int) (int, error)
}
