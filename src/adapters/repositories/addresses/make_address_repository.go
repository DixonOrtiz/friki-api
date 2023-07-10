package addressrepo

import "cloud.google.com/go/firestore"

func MakeAddressRepository(DB *firestore.Client) IAddressRepository {
	return &AddressRepository{
		DB: DB,
	}
}
