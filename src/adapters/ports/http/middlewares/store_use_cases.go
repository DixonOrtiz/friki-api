package midports

import "frikiapi/src/entities"

type StoreUseCases interface {
	GetByExternalID(externalID string) (entities.Store, error)
}
