package types

import (
	authinfra "frikiapi/src/infraestructure/auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Infraestructure struct {
	Router          *gin.Engine
	ProtectedRouter *gin.RouterGroup
	AuthConfig      authinfra.OAuth2ConfigInterface
	DB              *gorm.DB
}
