package main

import (
	"log"
	"os"

	"frikiapi/src/infraestructure/assembler"
	gormdb "frikiapi/src/infraestructure/db/gorm"
	oauthinfra "frikiapi/src/infraestructure/oauth"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	DB, err := gormdb.CreateConnection()
	if err != nil {
		log.Fatal(err)
	}

	assembler := assembler.MakeAssembler(
		oauthinfra.SetupConfig(),
		DB,
		gin.Default(),
	)
	router := assembler.GetRouterConfigured()
	router.Run(os.Getenv("PORT"))
}
