package addressusecases

import "frikiapi/src/entities"

type IAddressUseCases interface {
	Create(address entities.Address) (entities.Address, error)
	GetByID(tokenUserID, userID string, resources map[string]string) error
}
