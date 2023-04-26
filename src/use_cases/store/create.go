package storeusecases

import (
	cutils "frikiapi/src/use_cases/store/utils/create"
	"frikiapi/src/utils/consts"
	"frikiapi/src/utils/errors"
)

func (u StoreUseCases) Create(input cutils.CreateInput) (int, error) {
	addressID, err := u.AddressRepository.Create(input.Address)
	if err != nil {
		return 0, errors.NewError(consts.INTERNAL, err)
	}

	storeID, err := u.StoreRepository.Create(input.Store, addressID)
	if err != nil {
		return 0, errors.NewError(consts.INTERNAL, err)
	}

	return storeID, nil
}
