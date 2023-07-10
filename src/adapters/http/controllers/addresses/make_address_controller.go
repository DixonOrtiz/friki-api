package addresshttp

import (
	addressusecases "frikiapi/src/use_cases/addresses"
)

func MakeAddressControllers(addressUseCases addressusecases.IAddressUseCases) AddressControllers {
	return AddressControllers{
		AddressUseCases: addressUseCases,
	}

}
