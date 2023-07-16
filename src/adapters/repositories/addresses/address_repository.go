package addressrepo

import "cloud.google.com/go/firestore"

var ADDRESSES = "addresses"

type AddressRepository struct {
	DB *firestore.Client
}
