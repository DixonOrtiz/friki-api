package userusecases

import (
	"frikiapi/src/entities"
	"frikiapi/src/utils/errors"
	"frikiapi/src/utils/permissions"
	"time"

	"frikiapi/src/utils/uuid"
)

func (u UserUseCases) Create(user entities.User) (string, bool, error) {
	foundUser, _, err := u.UserRepository.GetByExternalID(user.ExternalID)
	if err != nil {
		return "", false, errors.New(errors.INTERNAL, err)
	}

	if foundUser.ID != "" {
		return foundUser.ID, false, nil
	}

	now := time.Now()

	user.ID = uuid.New()
	user.Role = permissions.USER
	user.UpdatedAt = now
	user.CreatedAt = now

	err = u.UserRepository.Create(user)
	if err != nil {
		return "", false, errors.New(errors.INTERNAL, err)
	}

	_, err = u.PermissionUseCases.Create(user.ID)
	if err != nil {
		return user.ID, true, errors.New(errors.INTERNAL, err)
	}

	return user.ID, true, nil
}
