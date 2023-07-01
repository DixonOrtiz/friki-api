package types

import (
	oauthinfra "frikiapi/src/infraestructure/oauth"

	"github.com/gin-gonic/gin"
)

type Infraestructure struct {
	Router          *gin.Engine
	ProtectedRouter *gin.RouterGroup
	OAuthConfig     oauthinfra.OAuth2ConfigInterface
}
