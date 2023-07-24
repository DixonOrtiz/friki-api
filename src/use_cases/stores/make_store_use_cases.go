package storeusecases

import (
	storerepository "frikiapi/src/adapters/repositories/stores"
	permusecases "frikiapi/src/use_cases/permissions"
)

func MakeStoreUseCases(
	storeRepository storerepository.IStoreRepository,
	permissionUseCases permusecases.IPermissionUseCases,
) IStoreUseCases {
	return &StoreUseCases{
		StoreRepository:    storeRepository,
		PermissionUseCases: permissionUseCases,
	}
}
