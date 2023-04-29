package storeusecases

import (
	"fmt"
	"frikiapi/src/utils/consts"
	"frikiapi/src/utils/errors"
)

func (u StoreUseCases) Authorize(externalID string, storeID int) error {
	store, err := u.StoreRepository.GetByExternalID(externalID)
	if err != nil {
		return errors.NewError(consts.INTERNAL, err)
	}

	if store.ID == 0 {
		return errors.NewError(consts.NOT_FOUND, fmt.Sprintf(
			"the store with external id %s does not exist",
			externalID,
		))
	}

	if store.ID != storeID {
		return errors.NewError(consts.UNAUTHORIZED, "only own stores can be modified")
	}

	return nil
}
