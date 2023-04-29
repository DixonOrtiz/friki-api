package storeusecases

import (
	"fmt"
	"frikiapi/src/entities"
	"frikiapi/src/utils/consts"
	"frikiapi/src/utils/errors"
)

func (u StoreUseCases) Create(store entities.Store) (int, error) {
	externalID := store.ExternalID

	exist, err := u.DoesExist(externalID, 0)
	if err != nil {
		return 0, errors.NewError(consts.INTERNAL, err)
	}

	if exist {
		return 0, errors.NewError(consts.UNPROCESSABLE, fmt.Sprintf(
			"external id %s is related to another store",
			externalID,
		))
	}

	storeID, err := u.StoreRepository.Create(store)
	if err != nil {
		return 0, errors.NewError(consts.INTERNAL, err)
	}

	return storeID, nil
}
