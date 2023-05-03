package oauthinfra

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	GOOGLE_CLIENT_ID     = "GOOGLE_CLIENT_ID"
	GOOGLE_CLIENT_SECRET = "GOOGLE_CLIENT_SECRET"
	REDIRECT_URL         = "REDIRECT_URL"
)

func SetupConfig() OAuth2ConfigInterface {
	return &oauth2.Config{
		ClientID:     os.Getenv(GOOGLE_CLIENT_ID),
		ClientSecret: os.Getenv(GOOGLE_CLIENT_SECRET),
		RedirectURL:  os.Getenv(REDIRECT_URL),
		Endpoint:     google.Endpoint,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
	}
}
