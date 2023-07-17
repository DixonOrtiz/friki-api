package addressusecases

import (
	"fmt"
	"frikiapi/src/entities"
	"frikiapi/src/utils/errors"
	"time"
)

func (u AddressUseCases) Update(address entities.Address) error {
	foundAddress, document, err := u.AddressRepository.GetByID(address.ID)
	if err != nil {
		return errors.New(errors.INTERNAL, err)
	}

	if foundAddress.ID == "" {
		return errors.New(errors.NOT_FOUND, fmt.Sprintf(
			"address with id '%s' is not in the registers",
			address.ID,
		))
	}

	address.UpdatedAt = time.Now()
	address.CreatedAt = foundAddress.CreatedAt

	err = u.AddressRepository.Update(document, address)
	if err != nil {
		return errors.New(errors.INTERNAL, err)
	}

	return nil
}
