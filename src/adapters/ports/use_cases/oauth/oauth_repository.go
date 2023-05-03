package authports

type OAuthRepository interface {
	GenerateExternalToken(code string) (string, error)
	GenerateLoginURL() string
}