package oauthrepo

type IOAuthRepository interface {
	GenerateExternalToken(code string) (string, error)
	GenerateLoginURL() string
}
