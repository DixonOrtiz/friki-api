package types

import (
	"frikiapi/src/entities"
)

func MapPermissionToEntity(firestorePermission FirestorePermission) entities.Permission {
	return entities.Permission{
		ID:        firestorePermission.ID,
		UserID:    firestorePermission.UserID,
		Addresses: firestorePermission.Addresses,
		Stores:    firestorePermission.Stores,
	}
}
