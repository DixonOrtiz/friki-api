package permusecases

import (
	"frikiapi/src/entities"
	"frikiapi/src/utils/errors"

	"frikiapi/src/utils/uuid"
)

func (u PermissionUseCases) Create(userID string) (entities.Permission, error) {
	permission := entities.Permission{
		ID:        uuid.New(),
		UserID:    userID,
		Addresses: []string{},
		Stores:    []string{},
	}

	err := u.PermissionRepository.Create(
		permission,
	)
	if err != nil {
		return entities.Permission{}, errors.New(errors.INTERNAL, err)
	}

	return permission, nil
}
