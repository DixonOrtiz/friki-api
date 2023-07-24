package storehttp

import storeusecases "frikiapi/src/use_cases/stores"

func MakeStoreControllers(storeUseCases storeusecases.IStoreUseCases) StoreControllers {
	return StoreControllers{
		StoreUseCases: storeUseCases,
	}
}
