package userusecases

import (
	"frikiapi/src/entities"
	"frikiapi/src/utils/consts"
	"frikiapi/src/utils/errors"
)

func (u UserUseCases) GetExternalUserByToken(token string) (entities.User, error) {
	user, err := u.UserRepository.GetExternalUserByToken(token)
	if err != nil {
		return entities.User{}, errors.New(consts.INTERNAL, err)
	}

	return user, nil

}
