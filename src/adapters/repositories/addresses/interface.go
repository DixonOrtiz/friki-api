package addressrepo

import "frikiapi/src/entities"

type IAddressRepository interface {
	Create(address entities.Address) error
}
