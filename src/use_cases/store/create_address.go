package storeusecases

import (
	"fmt"
	"frikiapi/src/entities"
	"frikiapi/src/utils/consts"
	"frikiapi/src/utils/errors"
)

func (u StoreUseCases) CreateAddress(address entities.Address) (int, error) {
	exist, err := u.DoesExist("", address.StoreID)
	if err != nil {
		return 0, errors.NewError(consts.INTERNAL, err)
	}

	if !exist {
		return 0, errors.NewError(consts.UNPROCESSABLE, fmt.Sprintf(
			"store with id %d does not exist",
			address.StoreID,
		))
	}

	addressID, err := u.AddressRepository.Create(address)
	if err != nil {
		return 0, errors.NewError(consts.INTERNAL, err)
	}

	return addressID, nil
}
