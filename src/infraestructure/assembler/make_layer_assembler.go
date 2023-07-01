package assembler

import (
	"frikiapi/src/infraestructure/assembler/types"
	oauthinfra "frikiapi/src/infraestructure/oauth"

	"github.com/gin-gonic/gin"
)

func MakeAssembler(
	oAuthConfig oauthinfra.OAuth2ConfigInterface,
	router *gin.Engine,
) Assembler {
	return Assembler{
		infraestructure: types.Infraestructure{
			OAuthConfig: oAuthConfig,
			Router:      router,
		},
	}
}
