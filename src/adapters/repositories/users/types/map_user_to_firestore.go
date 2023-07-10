package types

import (
	"frikiapi/src/entities"
)

func MapUserToFirestore(user entities.User) FirestoreUser {
	return FirestoreUser{
		ID:         user.ID,
		ExternalID: user.ExternalID,
		Name:       user.Name,
		LastName:   user.LastName,
		Email:      user.Email,
		Picture:    user.Picture,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}
}
