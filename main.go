package main

import (
	"log"
	"os"

	gormdb "frikiapi/src/infraestructure/db/gorm"
	oauthinfra "frikiapi/src/infraestructure/oauth"
	assembler "frikiapi/src/utils/layer_assembler"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	DB, err := gormdb.CreateConnection()
	if err != nil {
		log.Fatal(err)
	}

	layerAssembler := assembler.MakeLayerAssembler(
		oauthinfra.SetupConfig(),
		DB,
		gin.Default(),
	)
	router := layerAssembler.GetRouterConfigured()
	router.Run(os.Getenv("PORT"))
}
