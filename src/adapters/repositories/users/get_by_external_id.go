package userrepo

import (
	"context"
	"frikiapi/src/adapters/repositories/users/types"
	"frikiapi/src/entities"

	"google.golang.org/api/iterator"
)

func (r *UserRepository) GetByExternalID(externalID string) (entities.User, string, error) {
	var firestoreUser types.FirestoreUser
	var document string
	query := r.DB.Collection("users").Where("external_id", "==", externalID).Limit(1)

	iter := query.Documents(context.Background())
	defer iter.Stop()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return entities.User{}, "", err
		}
		doc.DataTo(&firestoreUser)
		document = doc.Ref.ID

	}

	return types.MapUserFirestoreToEntity(firestoreUser), document, nil
}
