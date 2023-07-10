package addressusecases

import addressrepository "frikiapi/src/adapters/repositories/addresses"

type AddressUseCases struct {
	AddressRepository addressrepository.IAddressRepository
}
