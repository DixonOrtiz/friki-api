package addressrepo

import (
	"context"
	"frikiapi/src/adapters/repositories/addresses/types"
	"frikiapi/src/entities"
)

func (r *AddressRepository) Update(document string, address entities.Address) error {
	firestoreAddress := types.MapAddressToFirestore(address)

	_, err := r.DB.Collection(ADDRESSES).Doc(document).Set(
		context.Background(),
		firestoreAddress,
	)
	if err != nil {
		return err
	}

	return nil
}
