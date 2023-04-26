package storeusecases

import (
	"fmt"
	"frikiapi/src/entities"
	cutils "frikiapi/src/use_cases/store/utils/create"
	"frikiapi/src/utils/consts"
	"frikiapi/src/utils/errors"
)

func (u StoreUseCases) Create(input cutils.CreateInput) (int, error) {
	externalID := input.Store.ExternalID

	store, err := u.StoreRepository.GetByExternalID(externalID)
	if err != nil {
		return 0, errors.NewError(consts.INTERNAL, err)
	}

	if store != (entities.Store{}) {
		return 0, errors.NewError(consts.UNPROCESSABLE, fmt.Sprintf(
			"external id %s is already in use",
			externalID,
		))
	}

	addressID, err := u.AddressRepository.Create(input.Address)
	if err != nil {
		return 0, errors.NewError(consts.INTERNAL, err)
	}

	input.Store.AddressID = addressID

	storeID, err := u.StoreRepository.Create(input.Store)
	if err != nil {
		u.AddressRepository.HardDeleteByID(addressID)
		return 0, errors.NewError(consts.INTERNAL, err)
	}

	return storeID, nil
}
