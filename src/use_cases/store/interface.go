package storeusecases

import cutils "frikiapi/src/use_cases/store/utils/create"

type StoreUseCasesInterface interface {
	Create(input cutils.CreateInput) (int, error)
}
