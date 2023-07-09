package userusecases

import (
	"frikiapi/src/entities"
	"frikiapi/src/utils/consts"
	"frikiapi/src/utils/errors"
)

func (u UserUseCases) Create(user entities.User) (bool, error) {
	foundUser, _, err := u.UserRepository.GetByExternalID(user.ExternalID)
	if err != nil {
		return false, errors.New(consts.Errors.INTERNAL, err)
	}

	if foundUser.ExternalID != "" {
		return false, nil
	}

	err = u.UserRepository.Create(user)
	if err != nil {
		return false, errors.New(consts.Errors.INTERNAL, err)
	}

	return true, nil
}
