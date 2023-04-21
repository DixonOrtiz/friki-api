package authrepo

func (r AuthRepository) GenerateLoginURL() string {
	return r.Config.AuthCodeURL("random_state")
}
