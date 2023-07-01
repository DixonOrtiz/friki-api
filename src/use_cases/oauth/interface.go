package authusecases

type IOAuthUseCases interface {
	GenerateLoginURL() string
	Login(code string) (string, bool, error)
}
