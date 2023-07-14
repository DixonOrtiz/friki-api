package permrepo

import "frikiapi/src/entities"

type IPermissionRepository interface {
	Create(permissions entities.Permission) error
	GetByUserID(userID string) (entities.Permission, string, error)
	AddResource(document string, permission entities.Permission) error
}
