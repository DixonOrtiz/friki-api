package userusecases

import (
	"fmt"
	"frikiapi/src/entities"
	"frikiapi/src/utils/errors"
)

func (u UserUseCases) GetByID(ID string) (entities.User, error) {
	user, _, err := u.UserRepository.GetByID(ID)
	if err != nil {
		return entities.User{}, errors.New(errors.INTERNAL, err)
	}

	if user.ID == "" {
		return entities.User{}, errors.New(
			errors.NOT_FOUND,
			fmt.Sprintf("user with id '%s' does not exist", ID),
		)
	}

	return user, nil
}
