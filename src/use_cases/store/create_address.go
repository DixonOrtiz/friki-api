package storeusecases

import (
	"frikiapi/src/entities"
	"frikiapi/src/utils/consts"
	"frikiapi/src/utils/errors"
)

func (u StoreUseCases) CreateAddress(address entities.Address) (int, error) {
	addressID, err := u.AddressRepository.Create(address)
	if err != nil {
		return 0, errors.NewError(consts.INTERNAL, err)
	}

	return addressID, nil
}
