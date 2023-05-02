package storeusecases

import (
	"frikiapi/src/entities"
	"frikiapi/src/utils/consts"
	"frikiapi/src/utils/errors"
)

func (u StoreUseCases) Update(store entities.Store) error {
	err := u.StoreRepository.Update(store)
	if err != nil {
		return errors.NewError(consts.INTERNAL, err)
	}

	return nil
}
