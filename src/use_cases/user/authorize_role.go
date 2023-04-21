package userusecases

import (
	"fmt"
	"frikiapi/src/entities"
	"frikiapi/src/utils"
	"frikiapi/src/utils/consts"
	"frikiapi/src/utils/errors"
)

func (u UserUseCases) AuthorizeRole(userID int, allowedRoles []int) (entities.User, error) {
	emptyUser := entities.User{}

	user, err := u.UserRepository.GetByID(userID)
	if err != nil {
		return emptyUser, errors.NewError(consts.INTERNAL, err)
	}

	if utils.DoesNumberExist(user.RoleID, allowedRoles) {
		return user, nil
	}

	roles := "roles"
	if len(allowedRoles) == 1 {
		roles = "role"
	}

	return emptyUser, errors.NewError(consts.UNAUTHORIZED, fmt.Sprintf(
		"this operation is allowed only for %s %s, current role is %d",
		roles,
		rolesToString(allowedRoles),
		user.RoleID,
	))
}

func rolesToString(roles []int) string {
	message := ""
	for i, role := range roles {
		if i == len(roles)-1 {
			message += fmt.Sprintf("%d", role)
			break
		}

		if i == len(roles)-2 {
			message += fmt.Sprintf("%d and ", role)
			continue
		}
		message = fmt.Sprintf("%d, ", role)
	}
	return message
}
