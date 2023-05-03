package oauthrepo

type OAuthRepositoryInterface interface {
	GenerateExternalToken(code string) (string, error)
	GenerateLoginURL() string
}
