package addressrepo

import (
	"context"
	"frikiapi/src/adapters/repositories/addresses/types"
	"frikiapi/src/entities"

	"google.golang.org/api/iterator"
)

func (r *AddressRepository) GetByID(ID string) (entities.Address, string, error) {
	var firestoreAddress types.FirestoreAddress
	var document string
	query := r.DB.Collection(ADDRESSES).Where("id", "==", ID).Limit(1)

	iter := query.Documents(context.Background())
	defer iter.Stop()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return entities.Address{}, "", err
		}
		doc.DataTo(&firestoreAddress)
		document = doc.Ref.ID

	}

	return types.MapUserFirestoreToEntity(firestoreAddress), document, nil
}
