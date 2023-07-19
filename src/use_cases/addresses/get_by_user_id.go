package addressusecases

import (
	"fmt"
	"frikiapi/src/entities"
	"frikiapi/src/utils/errors"
)

func (u AddressUseCases) GetByUserID(userID string) ([]entities.Address, error) {
	addresses, err := u.AddressRepository.GetByUserID(userID)
	if err != nil {
		return nil, errors.New(errors.INTERNAL, err)
	}

	if len(addresses) == 0 {
		return nil, errors.New(errors.NOT_FOUND, fmt.Sprintf("no address found with user_id %s", userID))
	}

	return addresses, nil
}
