package permrepo

import "frikiapi/src/entities"

type IPermissionRepository interface {
	Create(permissions entities.Permission) error
	GetByUserID(userID string) (entities.Permission, string, error)
	UpdateResource(document string, permission entities.Permission) error
}
