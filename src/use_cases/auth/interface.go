package authusecases

type AuthUseCasesInterface interface {
	GenerateLoginURL() string
	Login(code string) (string, bool, error)
}
