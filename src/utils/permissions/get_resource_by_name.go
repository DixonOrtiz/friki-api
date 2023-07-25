package permissions

import "frikiapi/src/entities"

func GetResourceByName(resource string, permission entities.Permission) []string {
	switch resource {
	case ADDRESS:
		return permission.Addresses
	case STORE:
		return permission.Stores
	}

	return nil
}
