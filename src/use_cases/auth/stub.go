package authusecases

import (
	"errors"
	"frikiapi/src/entities"
	jwtauth "frikiapi/src/utils/jwt"

	"github.com/undefinedlabs/go-mpatch"
)

var testExternalToken = "test_external_token"
var testURLInput = "google_login_url.example.com"
var expectedTestURL = "google_login_url.example.com"
var expectedExternalID = "AS17-AS7-AS10-AS70"

var testExternalUser = entities.User{
	Name:       "Alexis Alejandro",
	LastName:   "Sánchez Sánchez",
	Email:      "alexis.sanchez@gmail.com",
	Picture:    "https://abolaenossa.files.wordpress.com/2007/12/alexis.jpg",
	ExternalID: "AS17-AS7-AS10-AS70",
}

var testUser = entities.User{
	ID:         17,
	Name:       "Alexis Alejandro",
	LastName:   "Sánchez Sánchez",
	Email:      "alexis.sanchez@gmail.com",
	Cellphone:  "+56917107071",
	Rut:        "16927586-6",
	ExternalID: "AS17-AS7-AS10-AS70",
	Picture:    "https://abolaenossa.files.wordpress.com/2007/12/alexis.jpg",
}

func patchGenerateJWT(testError string) {
	mpatch.PatchMethod(jwtauth.GenerateJWT, func(externalID string) (string, error) {
		return "", errors.New(testError)
	})
}
