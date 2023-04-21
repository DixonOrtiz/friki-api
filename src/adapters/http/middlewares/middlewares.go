package middlewares

import midports "frikiapi/src/adapters/ports/http/middlewares"

type Middlewares struct {
	UserUseCases midports.UserUseCases
}
