package userrepo

import (
	"frikiapi/src/entities"
)

func (r *UserRepository) GetByExternalID(externalID string) (entities.User, error) {
	var user entities.User
	emptyUser := entities.User{}

	err := r.DB.Table("users").Where("external_id = ?", externalID).First(&user).Error

	if err != nil {
		if err.Error() == "record not found" {
			return emptyUser, nil
		}
		return emptyUser, err
	}

	if user.ID != 0 {
		return user, nil
	}

	return emptyUser, nil
}
