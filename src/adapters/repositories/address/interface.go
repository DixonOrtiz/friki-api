package addressrepo

import "frikiapi/src/entities"

type AddressRepositoryInterface interface {
	Create(address entities.Address) (int, error)
}
