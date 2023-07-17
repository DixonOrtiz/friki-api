package permusecases

import "frikiapi/src/entities"

type IPermissionUseCases interface {
	Authorize(JWTUserID string, userID string, resources map[string]string) error
	Create(userID string) (entities.Permission, error)
	AddResource(resource string, userID string, resourceID string) error
	RemoveResource(resource string, userID string, resourceID string) error
}
