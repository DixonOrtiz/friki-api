package middlewares

import midports "frikiapi/src/adapters/ports/http/middlewares"

func MakeMiddlewares(storeUseCases midports.StoreUseCases) Middlewares {
	return Middlewares{
		StoreUseCases: storeUseCases,
	}
}
