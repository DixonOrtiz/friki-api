package oauthrepo

func (r OAuthRepository) GenerateLoginURL() string {
	return r.Config.AuthCodeURL("random_state")
}
