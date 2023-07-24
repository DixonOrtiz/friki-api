package storeusecases

import (
	storerepository "frikiapi/src/adapters/repositories/stores"
	permusecases "frikiapi/src/use_cases/permissions"
)

type StoreUseCases struct {
	StoreRepository    storerepository.IStoreRepository
	PermissionUseCases permusecases.IPermissionUseCases
}
