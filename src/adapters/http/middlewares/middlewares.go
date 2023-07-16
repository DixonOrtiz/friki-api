package middlewares

import permusecases "frikiapi/src/use_cases/permissions"

type Middlewares struct {
	PermissionUseCases permusecases.IPermissionUseCases
}
