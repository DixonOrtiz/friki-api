package oauthrepo

import (
	oauthinfra "frikiapi/src/infraestructure/oauth"
)

func MakeOAuthRepository(config oauthinfra.OAuth2ConfigInterface) OAuthRepositoryInterface {
	return OAuthRepository{
		Config: config,
	}
}
