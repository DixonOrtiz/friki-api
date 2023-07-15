package addressusecases

import (
	"frikiapi/src/entities"
	"frikiapi/src/utils/errors"
	"frikiapi/src/utils/permissions"
	"time"

	"github.com/google/uuid"
)

func (u AddressUseCases) Create(address entities.Address) (entities.Address, error) {
	uid := uuid.New()
	now := time.Now()

	address.ID = uid.String()
	address.CreatedAt = now
	address.UpdatedAt = now

	err := u.AddressRepository.Create(address)
	if err != nil {
		return entities.Address{}, errors.New(errors.INTERNAL, err)
	}

	err = u.PermissionUseCases.AddResource(
		permissions.ADDRESS_ID,
		address.UserID,
		address.ID,
	)
	if err != nil {
		return entities.Address{}, errors.New(errors.INTERNAL, err)
	}

	return address, nil
}
