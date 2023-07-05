package userrepo

import (
	"context"
	"frikiapi/src/adapters/repositories/user/types"
	"frikiapi/src/entities"

	"google.golang.org/api/iterator"
)

func (r *UserRepository) GetByExternalID(externalID string) (entities.User, error) {
	var firestoreUser types.FirestoreUser
	query := r.DB.Collection("users").Where("external_id", "==", externalID).Limit(1)

	iter := query.Documents(context.Background())
	defer iter.Stop()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return entities.User{}, nil
		}
		doc.DataTo(&firestoreUser)
	}

	return types.MapUserFirestoreToEntity(firestoreUser), nil
}
