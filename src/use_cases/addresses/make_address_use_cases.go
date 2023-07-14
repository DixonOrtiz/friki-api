package addressusecases

import (
	addressrepository "frikiapi/src/adapters/repositories/addresses"
	permusecases "frikiapi/src/use_cases/permissions"
)

func MakeAddressUseCases(
	addressRepository addressrepository.IAddressRepository,
	permissionUseCases permusecases.IPermissionUseCases,
) IAddressUseCases {
	return &AddressUseCases{
		AddressRepository:  addressRepository,
		PermissionUseCases: permissionUseCases,
	}
}
