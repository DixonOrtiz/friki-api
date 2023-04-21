package authrepo

type AuthRepositoryInterface interface {
	GenerateExternalToken(code string) (string, error)
	GenerateLoginURL() string
}
