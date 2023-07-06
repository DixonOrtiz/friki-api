package userusecases

import (
	"fmt"
	"frikiapi/src/entities"
	"frikiapi/src/utils/consts"
	"frikiapi/src/utils/errors"
)

func (u UserUseCases) Update(user entities.User) error {
	exist, document, err := u.DoesExist(user.ExternalID)
	if err != nil {
		return errors.New(consts.INTERNAL, err)
	}

	if !exist {
		return errors.New(consts.NOT_FOUND, fmt.Sprintf(
			"user with external_id '%s' is not in the registers",
			user.ExternalID,
		))
	}

	err = u.UserRepository.Update(document, user)
	if err != nil {
		return errors.New(consts.INTERNAL, err)
	}

	return nil
}
