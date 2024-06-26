package addressrepo

import "frikiapi/src/entities"

type IAddressRepository interface {
	Create(address entities.Address) error
	GetByID(ID string) (entities.Address, string, error)
	GetByUserID(userID string) ([]entities.Address, error)
	Update(document string, address entities.Address) error
	Delete(document string) error
}
