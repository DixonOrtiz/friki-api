package storeusecases

import (
	storeports "frikiapi/src/adapters/ports/use_cases/store"
)

func MakeStoreUseCases(
	storeRepository storeports.StoreRepository,
	addressRepository storeports.AddressRepository,
) StoreUseCasesInterface {
	return &StoreUseCases{
		StoreRepository:   storeRepository,
		AddressRepository: addressRepository,
	}
}
