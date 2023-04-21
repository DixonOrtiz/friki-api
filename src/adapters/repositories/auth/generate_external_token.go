package authrepo

import "context"

func (r AuthRepository) GenerateExternalToken(code string) (string, error) {
	token, err := r.Config.Exchange(context.Background(), code)
	if err != nil {
		return "", err
	}

	return token.AccessToken, nil
}
