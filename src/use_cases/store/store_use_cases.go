package storeusecases

import storeports "frikiapi/src/adapters/ports/use_cases/store"

type StoreUseCases struct {
	StoreRepository   storeports.StoreRepository
	AddressRepository storeports.AddressRepository
}
