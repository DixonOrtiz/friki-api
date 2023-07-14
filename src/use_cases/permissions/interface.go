package permusecases

import "frikiapi/src/entities"

type IPermissionUseCases interface {
	Create(userID string) (entities.Permission, error)
	AddResource(resource string, userID string, resourceID string) error
}
