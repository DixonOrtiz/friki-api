package main

import (
	"fmt"
	"log"

	"frikiapi/src/infraestructure/assembler"
	"frikiapi/src/infraestructure/firestore"

	oauthinfra "frikiapi/src/infraestructure/oauth"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	DB, err := firestore.CreateConnection()
	if err != nil {
		log.Fatal(fmt.Sprintf("could not connect to firestore database. %e", err))
	}

	assembler := assembler.MakeAssembler(
		oauthinfra.SetupConfig(),
		gin.Default(),
		DB,
	)
	router := assembler.GetRouterConfigured()

	router.Run()
}
