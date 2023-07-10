package types

import (
	"frikiapi/src/entities"
)

func MapAddressToFirestore(address entities.Address) FirestoreAddress {
	return FirestoreAddress{
		ID:               address.ID,
		UserID:           address.UserID,
		Name:             address.Name,
		City:             address.City,
		Commune:          address.Commune,
		Street:           address.Street,
		Number:           address.Number,
		Apartment:        address.Apartment,
		Sector:           address.Sector,
		Type:             address.Type,
		ExtraInformation: address.ExtraInformation,
		CreatedAt:        address.CreatedAt,
		UpdatedAt:        address.UpdatedAt,
	}
}
