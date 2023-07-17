package addressusecases

import "frikiapi/src/entities"

type IAddressUseCases interface {
	Create(address entities.Address) (entities.Address, error)
	GetByID(addressID string) (entities.Address, error)
	Update(address entities.Address) error
	Delete(addressID string) error
}
