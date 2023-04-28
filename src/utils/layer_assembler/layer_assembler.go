package assembler

import (
	mid "frikiapi/src/adapters/http/middlewares"
	"frikiapi/src/utils/layer_assembler/types"
)

type LayerAssembler struct {
	middlewares     mid.Middlewares
	controllers     types.Controllers
	useCases        types.UseCases
	repositories    types.Repositories
	infraestructure types.Infraestructure
}
