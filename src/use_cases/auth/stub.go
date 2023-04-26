package authusecases

import (
	"errors"
	"frikiapi/src/entities"
	jwtauth "frikiapi/src/utils/jwt"

	"github.com/undefinedlabs/go-mpatch"
)

var (
	testExternalToken  = "test_external_token"
	testURLInput       = "google_login_url.example.com"
	expectedTestURL    = "google_login_url.example.com"
	expectedExternalID = "AS17-AS7-AS10-AS70"
)

var testExternalUser = entities.User{
	Name:       "Alexis Alejandro",
	LastName:   "Sánchez Sánchez",
	Email:      "alexis.sanchez@gmail.com",
	Picture:    "https://abolaenossa.files.wordpress.com/2007/12/alexis.jpg",
	ExternalID: "AS17-AS7-AS10-AS70",
}

func patchGenerateJWT(testError string) {
	mpatch.PatchMethod(jwtauth.GenerateJWT, func(externalID string) (string, error) {
		return "", errors.New(testError)
	})
}
