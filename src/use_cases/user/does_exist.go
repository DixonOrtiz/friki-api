package userusecases

import (
	"frikiapi/src/utils/consts"
	"frikiapi/src/utils/errors"
)

func (u UserUseCases) DoesExist(externalID string) (bool, error) {
	user, err := u.UserRepository.GetByExternalID(externalID)
	if err != nil {
		return false, errors.New(consts.INTERNAL, err)
	}

	if user.ExternalID != "" {
		return true, nil
	}

	return false, nil
}
