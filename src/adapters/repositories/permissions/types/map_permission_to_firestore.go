package types

import (
	"frikiapi/src/entities"
)

func MapPermissionToFirestore(permission entities.Permission) FirestorePermission {
	return FirestorePermission{
		ID:        permission.ID,
		UserID:    permission.UserID,
		Addresses: permission.Addresses,
		Stores:    permission.Stores,
	}
}
