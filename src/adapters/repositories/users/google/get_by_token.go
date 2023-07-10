package googleuserrepo

import (
	"encoding/json"
	"frikiapi/src/adapters/repositories/users/google/types"
	"frikiapi/src/entities"
	"io"
	"net/http"
)

func (r *GoogleUserRepository) GetByToken(token string) (entities.User, error) {
	emptyUser := entities.User{}
	usersURL := "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

	res, err := http.Get(usersURL + token)
	if err != nil {
		return emptyUser, err
	}

	defer res.Body.Close()

	buffer, err := io.ReadAll(res.Body)
	if err != nil {
		return emptyUser, err
	}

	var googleUser types.GoogleUser

	err = json.Unmarshal(buffer, &googleUser)
	if err != nil {
		return emptyUser, err
	}

	user := types.MapGoogleUserToEntity(googleUser)

	return user, nil
}
