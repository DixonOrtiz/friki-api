package types

import (
	"frikiapi/src/entities"
)

func MapUserFirestoreToEntity(firestoreUser FirestoreUser) entities.User {
	return entities.User{
		ExternalID: firestoreUser.ExternalID,
		Name:       firestoreUser.Name,
		LastName:   firestoreUser.LastName,
		Email:      firestoreUser.Email,
		Picture:    firestoreUser.Picture,
		CreatedAt:  firestoreUser.CreatedAt,
		UpdatedAt:  firestoreUser.UpdatedAt,
	}
}
