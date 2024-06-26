package permrepo

import (
	"context"
	"fmt"
	"frikiapi/src/adapters/repositories/permissions/types"
	"frikiapi/src/entities"

	"google.golang.org/api/iterator"
)

func (r *PermissionRepository) GetByUserID(userID string) (entities.Permission, string, error) {
	var firestorePermission types.FirestorePermission
	var document string

	query := r.DB.Collection("permissions").Where("user_id", "==", userID).Limit(1)

	iter := query.Documents(context.Background())
	defer iter.Stop()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return entities.Permission{}, "", err
		}
		doc.DataTo(&firestorePermission)
		document = doc.Ref.ID

	}

	if firestorePermission.ID == "" {
		return entities.Permission{}, "", fmt.Errorf(
			"the permission of user %s does not exist (check that there is no corrupt data)",
			userID,
		)
	}

	return types.MapPermissionToEntity(firestorePermission), document, nil
}
