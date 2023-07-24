package storerepo

import (
	"context"
	"frikiapi/src/adapters/repositories/stores/types"
	"frikiapi/src/entities"
)

func (r *StoreRepository) Create(store entities.Store) error {
	firestoreStore := types.MapStoreToFirestore(store)

	_, err := r.DB.Collection(STORES).NewDoc().Create(context.Background(), firestoreStore)
	if err != nil {
		return err
	}

	return nil
}
