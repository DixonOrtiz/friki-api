package permusecases

import (
	"fmt"
	"frikiapi/src/utils/errors"
	"frikiapi/src/utils/permissions"
)

func (u PermissionUseCases) AddResource(
	resource string,
	userID string,
	resourceID string,
) error {
	permission, document, err := u.PermissionRepository.GetByUserID(userID)
	if err != nil {
		errors.New(errors.INTERNAL, err)
	}

	switch resource {
	case permissions.ADDRESS:
		permission.Addresses = append(permission.Addresses, resourceID)
		u.PermissionRepository.AddResource(document, permission)

	default:
		return errors.New(errors.UNPROCESSABLE, fmt.Sprintf(
			"'%s' resource is not supported",
			resource,
		))
	}

	return nil
}
