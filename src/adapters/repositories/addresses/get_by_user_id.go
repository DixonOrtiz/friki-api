package addressrepo

import (
	"context"
	"frikiapi/src/adapters/repositories/addresses/types"
	"frikiapi/src/entities"

	"google.golang.org/api/iterator"
)

func (r *AddressRepository) GetByUserID(userID string) ([]entities.Address, error) {
	var addresses []entities.Address
	var firestoreAddress types.FirestoreAddress

	query := r.DB.Collection(ADDRESSES).Where("user_id", "==", userID)

	iter := query.Documents(context.Background())
	defer iter.Stop()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		doc.DataTo(&firestoreAddress)
		addresses = append(addresses, types.MapUserFirestoreToEntity(firestoreAddress))
	}

	return addresses, nil
}
