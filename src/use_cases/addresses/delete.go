package addressusecases

import (
	"fmt"
	"frikiapi/src/utils/errors"
)

func (u AddressUseCases) Delete(addressID string) error {
	foundAddress, document, err := u.AddressRepository.GetByID(addressID)
	if err != nil {
		return errors.New(errors.INTERNAL, err)
	}

	if foundAddress.ID == "" {
		return errors.New(errors.NOT_FOUND, fmt.Sprintf(
			"address with id '%s' is not in the registers",
			addressID,
		))
	}

	err = u.AddressRepository.Delete(document)
	if err != nil {
		return errors.New(errors.INTERNAL, err)
	}

	return nil
}
