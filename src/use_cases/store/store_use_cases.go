package storeusecases

import (
	storerepository "frikiapi/src/adapters/repositories/store"
)

type StoreUseCases struct {
	StoreRepository   storerepository.IStoreRepository
	
}
