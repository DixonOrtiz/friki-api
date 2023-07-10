package addressrepo

import (
	"context"
	"frikiapi/src/adapters/repositories/addresses/types"
	"frikiapi/src/entities"
)

func (r *AddressRepository) Create(address entities.Address) error {
	firestoreAddress := types.MapAddressToFirestore(address)

	_, err := r.DB.Collection("addresses").NewDoc().Create(context.Background(), firestoreAddress)
	if err != nil {
		return err
	}

	return nil
}
