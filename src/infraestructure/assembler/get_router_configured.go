package assembler

import (
	"github.com/gin-gonic/gin"
)

func (a *Assembler) GetRouterConfigured() *gin.Engine {
	a.setRepositories()
	a.setUseCases()
	a.setControllers()
	a.setRouters()

	return a.infraestructure.Router
}
