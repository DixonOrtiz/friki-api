package userrepo

import (
	"fmt"
	"frikiapi/src/entities"
)

func (r *UserRepository) Update(user entities.User) error {
	res := r.DB.
		Table("users").
		Omit("external_id").
		Updates(user).
		Where("id = ?", user.ID)

	err := res.Error
	if err != nil {
		return err
	}

	if res.RowsAffected == 0 {
		return fmt.Errorf("there is no user with id %d", user.ID)
	}

	return nil
}
