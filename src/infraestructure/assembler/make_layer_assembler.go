package assembler

import (
	"frikiapi/src/infraestructure/assembler/types"
	oauthinfra "frikiapi/src/infraestructure/oauth"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
)

func MakeAssembler(
	oAuthConfig oauthinfra.OAuth2ConfigInterface,
	router *gin.Engine,
	DB *firestore.Client,
) Assembler {
	return Assembler{
		infraestructure: types.Infraestructure{
			OAuthConfig: oAuthConfig,
			Router:      router,
			DB:          DB,
		},
	}
}
