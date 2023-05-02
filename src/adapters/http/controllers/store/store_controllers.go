package storehttp

import (
	contports "frikiapi/src/adapters/ports/http/controllers"
)

type StoreControllers struct {
	StoreUseCases contports.StoreUseCases
}
