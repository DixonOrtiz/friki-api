package main

import (
	"log"
	"os"

	"frikiapi/src/infraestructure/assembler"
	"frikiapi/src/infraestructure/firestore"

	oauthinfra "frikiapi/src/infraestructure/oauth"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	_, err := firestore.CreateConnection()
	if err != nil {
		log.Fatal("could not connect to firestore database")
	}

	assembler := assembler.MakeAssembler(
		oauthinfra.SetupConfig(),
		gin.Default(),
	)
	router := assembler.GetRouterConfigured()
	router.Run(os.Getenv("PORT"))
}
