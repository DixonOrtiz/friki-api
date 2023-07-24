package storeusecases

import "frikiapi/src/entities"

type IStoreUseCases interface {
	Create(store entities.Store) (entities.Store, error)
}
