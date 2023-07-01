package assembler

import (
	mid "frikiapi/src/adapters/http/middlewares"
	"frikiapi/src/infraestructure/assembler/types"
)

type Assembler struct {
	middlewares     mid.Middlewares
	controllers     types.Controllers
	useCases        types.UseCases
	repositories    types.Repositories
	infraestructure types.Infraestructure
}
