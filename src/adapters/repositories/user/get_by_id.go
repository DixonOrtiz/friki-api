package userrepo

import "frikiapi/src/entities"

func (r UserRepository) GetByID(userID int) (entities.User, error) {
	var user entities.User
	emptyUser := entities.User{}

	res := r.DB.Table("users").
		Where("id = ?", userID).
		Where("deleted_at IS NULL").
		Limit(1).
		First(&user)

	err := res.Error
	if err != nil {
		return emptyUser, err
	}

	if res.RowsAffected == 0 {
		return emptyUser, nil
	}

	return user, nil
}
