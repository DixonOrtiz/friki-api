package userusecases

import (
	"frikiapi/src/entities"
	"frikiapi/src/utils/consts"
	"frikiapi/src/utils/errors"
)

func (u UserUseCases) Update(user entities.User) error {
	err := u.UserRepository.Update(user)
	if err != nil {
		return errors.New(consts.INTERNAL, err)
	}

	return nil
}
