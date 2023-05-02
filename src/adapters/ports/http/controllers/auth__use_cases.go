package contports

type AuthUseCases interface {
	GenerateLoginURL() string
	Login(code string) (string, error)
}
