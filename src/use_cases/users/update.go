package userusecases

import (
	"fmt"
	"frikiapi/src/entities"
	"frikiapi/src/utils/errors"
	"time"
)

func (u UserUseCases) Update(user entities.User) error {
	foundUser, document, err := u.UserRepository.GetByID(user.ID)
	if err != nil {
		return errors.New(errors.INTERNAL, err)
	}

	if foundUser.ID == "" {
		return errors.New(errors.NOT_FOUND, fmt.Sprintf(
			"user with id '%s' is not in the registers",
			user.ID,
		))
	}

	user.ExternalID = foundUser.ExternalID
	user.Email = foundUser.Email
	user.UpdatedAt = time.Now()
	user.CreatedAt = foundUser.CreatedAt

	err = u.UserRepository.Update(document, user)
	if err != nil {
		return errors.New(errors.INTERNAL, err)
	}

	return nil
}
