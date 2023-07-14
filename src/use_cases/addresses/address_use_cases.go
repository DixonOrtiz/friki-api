package addressusecases

import (
	addressrepository "frikiapi/src/adapters/repositories/addresses"
	permusecases "frikiapi/src/use_cases/permissions"
)

type AddressUseCases struct {
	AddressRepository  addressrepository.IAddressRepository
	PermissionUseCases permusecases.IPermissionUseCases
}
