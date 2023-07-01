package storeusecases

import (
	addressepository "frikiapi/src/adapters/repositories/address"
	storerepository "frikiapi/src/adapters/repositories/store"
)

type StoreUseCases struct {
	StoreRepository   storerepository.IStoreRepository
	AddressRepository addressepository.IAddressRepository
}
