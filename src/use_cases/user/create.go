package userusecases

import (
	"frikiapi/src/entities"
	"frikiapi/src/utils/consts"
	"frikiapi/src/utils/errors"
)

func (u UserUseCases) Create(user entities.User) (bool, error) {
	exists, err := u.DoesExist(user.ExternalID)
	if err != nil {
		return false, errors.New(consts.INTERNAL, err)
	}

	if exists {
		return false, nil
	}

	err = u.UserRepository.Create(user)
	if err != nil {
		return false, errors.New(consts.INTERNAL, err)
	}

	return true, nil
}
