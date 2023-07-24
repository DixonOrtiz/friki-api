package types

import (
	"frikiapi/src/entities"
)

func MapFirestoreStoreToEntity(firestoreStore FirestoreStore) entities.Store {
	return entities.Store{
		ID:          firestoreStore.ID,
		UserID:      firestoreStore.UserID,
		Name:        firestoreStore.Name,
		Description: firestoreStore.Description,
		CreatedAt:   firestoreStore.CreatedAt,
		UpdatedAt:   firestoreStore.UpdatedAt,
	}
}
