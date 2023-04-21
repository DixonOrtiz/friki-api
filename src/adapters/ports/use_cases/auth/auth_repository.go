package authports

type AuthRepository interface {
	GenerateExternalToken(code string) (string, error)
	GenerateLoginURL() string
}
