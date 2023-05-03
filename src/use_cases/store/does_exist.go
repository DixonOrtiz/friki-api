package storeusecases

import (
	"frikiapi/src/entities"
	"frikiapi/src/utils/consts"
	"frikiapi/src/utils/errors"
)

func (u StoreUseCases) DoesExist(externalID string, ID int) (bool, error) {
	var (
		store entities.Store
		err   error
	)

	if externalID != "" {
		store, err = u.StoreRepository.GetByExternalID(externalID)
		if err != nil {
			return false, errors.New(consts.INTERNAL, err)
		}
	}

	if ID != 0 {
		store, err = u.StoreRepository.GetByID(ID)
		if err != nil {
			return false, errors.New(consts.INTERNAL, err)
		}
	}

	if store != (entities.Store{}) {
		return true, nil
	}

	return false, nil
}
