package oauthrepo

import (
	oauthinfra "frikiapi/src/infraestructure/oauth"
)

func MakeOAuthRepository(config oauthinfra.OAuth2ConfigInterface) IOAuthRepository {
	return OAuthRepository{
		Config: config,
	}
}
