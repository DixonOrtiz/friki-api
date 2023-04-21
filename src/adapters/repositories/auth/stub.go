package authrepo

import "golang.org/x/oauth2"

var expectedExternalToken = "test_external_token"
var emptyToken = &oauth2.Token{}
var testLoginURL = "google.com/users/url"

func getTestToken() *oauth2.Token {
	return &oauth2.Token{
		AccessToken: "test_external_token",
	}
}
