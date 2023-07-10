package permusecases

import (
	"frikiapi/src/entities"
	"frikiapi/src/utils/errors"

	"github.com/google/uuid"
)

func (u PermissionUseCases) Create(userID string) (entities.Permission, error) {
	uid := uuid.New()

	permission := entities.Permission{
		ID:        uid.String(),
		UserID:    userID,
		Addresses: []string{},
	}

	err := u.PermissionRepository.Create(
		permission,
	)
	if err != nil {
		return entities.Permission{}, errors.New(errors.INTERNAL, err)
	}

	return permission, nil
}
