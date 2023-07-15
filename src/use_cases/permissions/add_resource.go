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
		return errors.New(errors.INTERNAL, err)
	}

	switch resource {
	case permissions.ADDRESS_ID:
		permission.Addresses = append(permission.Addresses, resourceID)
		err = u.PermissionRepository.AddResource(document, permission)
		if err != nil {
			return errors.New(errors.INTERNAL, err)
		}

	default:
		return errors.New(errors.CONFLICT, fmt.Sprintf(
			"'%s' resource is not supported",
			resource,
		))
	}

	return nil
}
