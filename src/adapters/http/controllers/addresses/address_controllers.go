package addresshttp

import (
	addressusecases "frikiapi/src/use_cases/addresses"
)

type AddressControllers struct {
	AddressUseCases addressusecases.IAddressUseCases
}
