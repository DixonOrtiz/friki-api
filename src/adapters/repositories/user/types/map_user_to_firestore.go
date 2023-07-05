package types

import (
	"frikiapi/src/entities"
	"time"
)

func MapUserToFirestore(user entities.User) FirestoreUser {
	now := time.Now()
	return FirestoreUser{
		ExternalID: user.ExternalID,
		Name: user.Name,
		LastName: user.LastName,
		Email: user.Email,
		Picture: user.Picture,
		CreatedAt: now,
		UpdatedAt: now,
	}
}