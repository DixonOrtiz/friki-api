package userports

import "frikiapi/src/entities"

type AddressRepository interface {
	Create(address entities.Address) (int, error)
	HardDeleteByID(ID int)
}
