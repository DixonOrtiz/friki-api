package permrepo

import (
	"context"
	"frikiapi/src/adapters/repositories/permissions/types"
	"frikiapi/src/entities"
)

func (r *PermissionRepository) Create(permission entities.Permission) error {
	firestorePermission := types.MapPermissionToFirestore(permission)

	_, err := r.DB.Collection("permissions").NewDoc().Create(context.Background(), firestorePermission)
	if err != nil {
		return err
	}

	return nil
}
