package types

import (
	oauthinfra "frikiapi/src/infraestructure/oauth"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
)

type Infraestructure struct {
	Router          *gin.Engine
	ProtectedRouter *gin.RouterGroup
	OAuthConfig     oauthinfra.OAuth2ConfigInterface
	DB              *firestore.Client
}
