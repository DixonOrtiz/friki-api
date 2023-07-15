package permusecases

import (
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
		errors.New(errors.UNAUTHORIZED, "path user_id and token user_id are not the same")
	}

	if len(resources) > 0 {
		permission, _, err := u.PermissionRepository.GetByUserID(userID)
		if err != nil {
			errors.New(errors.INTERNAL, err)
		}

		for resource, resourceID := range resources {
			userResources := permissions.GetResourceByName(resource, permission)

			if !slices.Exist(userResources, resourceID) {
				return errors.New(errors.UNAUTHORIZED, "the requested resource is not owned by this user")
			}
		}

	}

	return nil
}
