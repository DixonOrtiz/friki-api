package types

import "frikiapi/src/entities"

func MapStoreToFirestore(store entities.Store) FirestoreStore {
	return FirestoreStore{
		ID:          store.ID,
		UserID:      store.UserID,
		Name:        store.Name,
		Description: store.Description,
		CreatedAt:   store.CreatedAt,
		UpdatedAt:   store.UpdatedAt,
	}
}
