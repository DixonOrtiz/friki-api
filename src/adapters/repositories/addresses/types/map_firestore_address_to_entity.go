package types

import (
	"frikiapi/src/entities"
)

func MapUserFirestoreToEntity(firestoreAddress FirestoreAddress) entities.Address {
	return entities.Address{
		ID:               firestoreAddress.ID,
		UserID:           firestoreAddress.UserID,
		Name:             firestoreAddress.Name,
		City:             firestoreAddress.City,
		Commune:          firestoreAddress.Commune,
		Street:           firestoreAddress.Street,
		Number:           firestoreAddress.Number,
		Apartment:        firestoreAddress.Apartment,
		Sector:           firestoreAddress.Sector,
		Type:             firestoreAddress.Type,
		ExtraInformation: firestoreAddress.ExtraInformation,
		CreatedAt:        firestoreAddress.CreatedAt,
		UpdatedAt:        firestoreAddress.UpdatedAt,
	}
}
