package oauthrepo

import oauthinfra "frikiapi/src/infraestructure/oauth"

type OAuthRepository struct {
	Config oauthinfra.OAuth2ConfigInterface
}
