package permusecases

import permrepository "frikiapi/src/adapters/repositories/permissions"

func MakePermissionUseCases(
	permissionRepository permrepository.IPermissionRepository,
) IPermissionUseCases {
	return &PermissionUseCases{
		PermissionRepository: permissionRepository,
	}
}
