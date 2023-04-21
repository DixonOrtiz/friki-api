package assembler

import (
	"frikiapi/src/utils/layer_assembler/types"
)

type LayerAssembler struct {
	controllers     types.Controllers
	useCases        types.UseCases
	repositories    types.Repositories
	infraestructure types.Infraestructure
}
