package userrepo

import "frikiapi/src/entities"

func (r *UserRepository) Create(user entities.User) (int, error) {
	err := r.DB.Table("users").Create(&user).Error
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}
