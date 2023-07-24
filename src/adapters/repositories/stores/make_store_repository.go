package storerepo

import "cloud.google.com/go/firestore"

func MakeStoreRepository(DB *firestore.Client) IStoreRepository {
	return &StoreRepository{
		DB: DB,
	}
}
