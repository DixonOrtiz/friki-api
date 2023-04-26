package userports

import "frikiapi/src/entities"

type StoreRepository interface {
	Create(store entities.Store, addressID int) (int, error)
}
