package storeports

import "frikiapi/src/entities"

type AddressRepository interface {
	Create(address entities.Address) (int, error)
}
