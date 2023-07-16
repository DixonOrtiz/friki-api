package permusecases

import (
	"fmt"
	"frikiapi/src/utils/errors"
	"frikiapi/src/utils/permissions"
	"frikiapi/src/utils/slices"
)

func (u PermissionUseCases) Authorize(
	JWTUserID string,
	userID string,
	resources map[string]string,
) error {
	if userID != JWTUserID {
		return errors.New(errors.UNAUTHORIZED, "path user_id and token user_id are not the same")
	}

	if len(resources) > 0 {
		permission, _, err := u.PermissionRepository.GetByUserID(userID)
		if err != nil {
			return errors.New(errors.INTERNAL, err)
		}

		for resource, ID := range resources {
			userResource := permissions.GetResourceByName(resource, permission)

			if !slices.Exist(userResource, ID) {
				return errors.New(
					errors.UNAUTHORIZED,
					fmt.Sprintf("the requested resource (%s) is not owned by this user", resource),
				)
			}
		}

	}

	return nil
}
