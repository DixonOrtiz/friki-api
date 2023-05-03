package authusecases

func (u OAuthUseCases) GenerateLoginURL() string {
	return u.OAuthRepository.GenerateLoginURL()
}
