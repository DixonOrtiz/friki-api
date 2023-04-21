package authusecases

func (u AuthUseCases) GenerateLoginURL() string {
	return u.AuthRepository.GenerateLoginURL()
}
