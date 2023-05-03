package oauthinfra

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	testGoogleClientID     = "123.apps.google.com"
	testGoogleClientSecret = "GSDL-1235-AD2Aa3jh2d"
	testRedirectUrl        = "http://localhost:8080/callback"
)

var expectedConfig = &oauth2.Config{
	ClientID:     "123.apps.google.com",
	ClientSecret: "GSDL-1235-AD2Aa3jh2d",
	RedirectURL:  "http://localhost:8080/callback",
	Endpoint:     google.Endpoint,
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
	},
}
