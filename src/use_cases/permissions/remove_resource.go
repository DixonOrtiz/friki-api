package permusecases

import (
	"fmt"
	"frikiapi/src/utils/errors"
	"frikiapi/src/utils/permissions"
	"frikiapi/src/utils/slices"
)

func (u PermissionUseCases) RemoveResource(
	resource string,
	userID string,
	resourceID string,
) error {
	permission, document, err := u.PermissionRepository.GetByUserID(userID)
	if err != nil {
		return errors.New(errors.INTERNAL, err)
	}

	switch resource {
	case permissions.ADDRESS:
		permission.Addresses = slices.RemoveElement(permission.Addresses, resourceID)
		err = u.PermissionRepository.UpdateResource(document, permission)

	case permissions.STORE:
		permission.Stores = slices.RemoveElement(permission.Stores, resourceID)
		err = u.PermissionRepository.UpdateResource(document, permission)

	default:
		return errors.New(errors.CONFLICT, fmt.Sprintf(
			"'%s' resource is not supported",
			resource,
		))
	}

	if err != nil {
		return errors.New(errors.INTERNAL, err)
	}

	return nil
}
