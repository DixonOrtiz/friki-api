package middlewares

import permusecases "frikiapi/src/use_cases/permissions"

func MakeMiddlewares(permissionUseCases permusecases.IPermissionUseCases) Middlewares {
	return Middlewares{
		PermissionUseCases: permissionUseCases,
	}
}
