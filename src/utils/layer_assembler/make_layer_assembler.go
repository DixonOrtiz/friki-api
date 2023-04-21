package assembler

import (
	authinfra "frikiapi/src/infraestructure/auth"
	"frikiapi/src/utils/layer_assembler/types"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MakeLayerAssembler(
	authConfig authinfra.OAuth2ConfigInterface,
	DB *gorm.DB,
	router *gin.Engine,
) LayerAssembler {
	return LayerAssembler{
		infraestructure: types.Infraestructure{
			AuthConfig: authConfig,
			DB:         DB,
			Router:     router,
		},
	}
}
