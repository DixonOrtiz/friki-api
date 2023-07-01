package storeusecases

import (
	storerepository "frikiapi/src/adapters/repositories/store"
)

func MakeStoreUseCases(
	storeRepository storerepository.IStoreRepository,
	
) IStoreUseCases {
	return &StoreUseCases{
		StoreRepository:   storeRepository,
	}
}
