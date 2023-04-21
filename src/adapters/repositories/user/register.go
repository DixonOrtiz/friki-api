package userrepo

import (
	"frikiapi/src/entities"
)

func (r UserRepository) Register(user entities.User) error {
	result := r.DB.Table("users").Create(&user)

	if err := result.Error; err != nil {
		return err
	}

	return nil
}
