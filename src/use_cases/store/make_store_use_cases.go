package storeusecases

import (
	storeports "frikiapi/src/adapters/ports/use_cases/store"
)

func MakeStoreUseCases(
	storeRepository storeports.StoreRepository,
) StoreUseCasesInterface {
	return &StoreUseCases{
		StoreRepository: storeRepository,
	}
}
