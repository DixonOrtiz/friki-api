package permrepo

import "frikiapi/src/entities"

type IPermissionRepository interface {
	Create(permissions entities.Permission) error
}
