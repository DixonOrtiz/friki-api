package userusecases

import (
	"frikiapi/src/entities"
	"frikiapi/src/utils/consts"
	"frikiapi/src/utils/errors"
)

func (u UserUseCases) GetByID(userID int) (entities.User, error) {
	user, err := u.UserRepository.GetByID(userID)
	if err != nil {
		return entities.User{}, errors.NewError(consts.INTERNAL, err)
	}

	return user, nil
}
