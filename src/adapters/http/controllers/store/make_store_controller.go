package storehttp

import storeusecases "frikiapi/src/use_cases/store"

func MakeStoreControllers(storeUseCases storeusecases.StoreUseCasesInterface) StoreControllers {
	return StoreControllers{
		StoreUseCases: storeUseCases,
	}

}
