package storerepo

import (
	"context"
	"frikiapi/src/adapters/repositories/stores/types"
	"frikiapi/src/entities"

	"google.golang.org/api/iterator"
)

func (r *StoreRepository) GetByID(ID string) (entities.Store, string, error) {
	var firestoreStore types.FirestoreStore
	var document string
	query := r.DB.Collection(STORES).Where("id", "==", ID).Limit(1)

	iter := query.Documents(context.Background())
	defer iter.Stop()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return entities.Store{}, "", err
		}
		doc.DataTo(&firestoreStore)
		document = doc.Ref.ID

	}

	return types.MapFirestoreStoreToEntity(firestoreStore), document, nil
}
