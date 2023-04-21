package assembler

import (
	"github.com/gin-gonic/gin"
)

func (a *LayerAssembler) GetRouterConfigured() *gin.Engine {
	a.setRepositories()
	a.setUseCases()
	a.setControllers()
	a.setRouters()

	return a.infraestructure.Router
}
