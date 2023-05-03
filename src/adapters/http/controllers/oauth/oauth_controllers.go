package oauthhttp

import (
	ports "frikiapi/src/adapters/ports/http/controllers/oauth"
)

type OAuthControllers struct {
	OAuthUseCases ports.OAuthUseCases
}
