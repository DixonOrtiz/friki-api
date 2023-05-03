package contports

type OAuthUseCases interface {
	GenerateLoginURL() string
	Login(code string) (string, bool, error)
}
