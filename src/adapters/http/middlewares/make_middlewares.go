package middlewares

import (
	midports "frikiapi/src/adapters/ports/http/middlewares"
)

func MakeMiddlewares(userUseCases midports.UserUseCases) Middlewares {
	return Middlewares{
		UserUseCases: userUseCases,
	}
}
