package addressrepo

import "cloud.google.com/go/firestore"

type AddressRepository struct {
	DB *firestore.Client
}
