package addressusecases

import (
	"fmt"
	"frikiapi/src/entities"
	"frikiapi/src/utils/errors"
)

func (u AddressUseCases) GetByID(addressID string) (entities.Address, error) {
	address, _, err := u.AddressRepository.GetByID(addressID)
	if err != nil {
		return entities.Address{}, errors.New(errors.INTERNAL, err)
	}

	if address.ID == "" {
		return entities.Address{}, errors.New(
			errors.NOT_FOUND,
			fmt.Sprintf("address with id %s not found", addressID),
		)
	}

	return address, nil
}
