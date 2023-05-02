package contports

import "frikiapi/src/entities"

type StoreUseCases interface {
	CreateAddress(address entities.Address) (int, error)
	Create(store entities.Store) (int, error)
	Update(store entities.Store) error
}
