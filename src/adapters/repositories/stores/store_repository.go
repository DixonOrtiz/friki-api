package storerepo

import "cloud.google.com/go/firestore"

var STORES = "stores"

type StoreRepository struct {
	DB *firestore.Client
}
