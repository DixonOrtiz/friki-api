package authusecases

import (
	"frikiapi/src/utils/consts"
	"frikiapi/src/utils/errors"
	jwtauth "frikiapi/src/utils/jwt"
)

func (u OAuthUseCases) Login(code string) (string, bool, error) {
	externalToken, err := u.OAuthRepository.GenerateExternalToken(code)
	if err != nil {
		return "", false, errors.New(consts.INTERNAL, err)
	}

	externalUser, err := u.UserUseCases.GetExternalUserByToken(externalToken)
	if err != nil {
		return "", false, err
	}

	created, err := u.UserUseCases.Create(externalUser)
	if err != nil {
		return "", false, errors.New(consts.INTERNAL, err)
	}

	token, err := jwtauth.GenerateJWT(externalUser.ExternalID)
	if err != nil {
		return "", false, errors.New(consts.INTERNAL, err)
	}

	return token, created, nil
}
