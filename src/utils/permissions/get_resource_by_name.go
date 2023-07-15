package permissions

import "frikiapi/src/entities"

func GetResourceByName(resource string, permission entities.Permission) []string {
	switch resource {
	case ADDRESS_ID:
		return permission.Addresses
	}

	return nil
}
