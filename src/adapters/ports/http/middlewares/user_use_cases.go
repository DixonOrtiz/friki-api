package midports

import "frikiapi/src/entities"

type UserUseCases interface {
	GetByExternalID(externalID string) (entities.User, error)
}
