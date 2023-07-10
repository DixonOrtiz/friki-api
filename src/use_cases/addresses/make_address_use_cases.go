package addressusecases

import addressrepository "frikiapi/src/adapters/repositories/addresses"

func MakeAddressUseCases(
	addressRepository addressrepository.IAddressRepository,
) IAddressUseCases {
	return &AddressUseCases{
		AddressRepository: addressRepository,
	}
}
