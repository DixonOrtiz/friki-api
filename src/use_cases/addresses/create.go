package addressusecases

import (
	"frikiapi/src/entities"
	"frikiapi/src/utils/errors"
	"frikiapi/src/utils/permissions"
	"frikiapi/src/utils/uuid"
	"time"
)

func (u AddressUseCases) Create(address entities.Address) (entities.Address, error) {
	now := time.Now()

	address.ID = uuid.New()
	address.CreatedAt = now
	address.UpdatedAt = now

	err := u.AddressRepository.Create(address)
	if err != nil {
		return entities.Address{}, errors.New(errors.INTERNAL, err)
	}

	err = u.PermissionUseCases.AddResource(
		permissions.ADDRESS,
		address.UserID,
		address.ID,
	)
	if err != nil {
		return entities.Address{}, errors.New(errors.INTERNAL, err)
	}

	return address, nil
}
