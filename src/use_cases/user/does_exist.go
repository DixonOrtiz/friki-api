package userusecases

import (
	"frikiapi/src/utils/consts"
	"frikiapi/src/utils/errors"
)

func (u UserUseCases) DoesExist(externalID string) (bool, string, error) {
	user, document, err := u.UserRepository.GetByExternalID(externalID)
	if err != nil {
		return false, "", errors.New(consts.INTERNAL, err)
	}

	if user.ExternalID != "" {
		return true, document, nil
	}

	return false, "", nil
}
