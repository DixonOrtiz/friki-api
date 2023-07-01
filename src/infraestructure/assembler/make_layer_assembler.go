package assembler

import (
	"frikiapi/src/infraestructure/assembler/types"
	oauthinfra "frikiapi/src/infraestructure/oauth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MakeAssembler(
	oAuthConfig oauthinfra.OAuth2ConfigInterface,
	DB *gorm.DB,
	router *gin.Engine,
) Assembler {
	return Assembler{
		infraestructure: types.Infraestructure{
			OAuthConfig: oAuthConfig,
			DB:          DB,
			Router:      router,
		},
	}
}
