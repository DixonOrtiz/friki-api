package main

import (
	"log"

	authinfra "frikiapi/src/infraestructure/auth"
	gormdb "frikiapi/src/infraestructure/db/gorm"
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
		authinfra.SetupConfig(),
		DB,
		gin.Default(),
	)
	router := layerAssembler.GetRouterConfigured()
	router.Run()
}
