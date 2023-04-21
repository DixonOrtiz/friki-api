package userusecases

import (
	"frikiapi/src/entities"
	"frikiapi/src/utils/consts"
	"frikiapi/src/utils/errors"
)

func (u UserUseCases) GetByExternalID(externalID string) (entities.User, error) {
	user, err := u.UserRepository.GetByExternalID(externalID)
	if err != nil {
		return entities.User{}, errors.NewError(consts.INTERNAL, err)
	}

	return user, nil
}
