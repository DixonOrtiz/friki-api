package addressrepo

import (
	"context"
)

func (r *AddressRepository) Delete(document string) error {
	_, err := r.DB.Collection(ADDRESSES).Doc(document).Delete(
		context.Background(),
	)
	if err != nil {
		return err
	}

	return nil
}
