package storeusecases

import (
	addressepository "frikiapi/src/adapters/repositories/address"
	storerepository "frikiapi/src/adapters/repositories/store"
)

func MakeStoreUseCases(
	storeRepository storerepository.IStoreRepository,
	addressRepository addressepository.IAddressRepository,
) IStoreUseCases {
	return &StoreUseCases{
		StoreRepository:   storeRepository,
		AddressRepository: addressRepository,
	}
}
