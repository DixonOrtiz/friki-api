package assembler

import (
	oauthinfra "frikiapi/src/infraestructure/oauth"
	"frikiapi/src/utils/layer_assembler/types"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MakeLayerAssembler(
	oAuthConfig oauthinfra.OAuth2ConfigInterface,
	DB *gorm.DB,
	router *gin.Engine,
) LayerAssembler {
	return LayerAssembler{
		infraestructure: types.Infraestructure{
			OAuthConfig: oAuthConfig,
			DB:          DB,
			Router:      router,
		},
	}
}
