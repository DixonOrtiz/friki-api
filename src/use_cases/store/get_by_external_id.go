package storeusecases

import (
	"frikiapi/src/entities"
	"frikiapi/src/utils/consts"
	"frikiapi/src/utils/errors"
)

func (u StoreUseCases) GetByExternalID(externalID string) (entities.Store, error) {
	store, err := u.StoreRepository.GetByExternalID(externalID)
	if err != nil {
		return store, errors.NewError(consts.INTERNAL, err)
	}

	return store, nil
}
