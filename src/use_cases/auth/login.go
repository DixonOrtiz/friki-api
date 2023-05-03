package authusecases

import (
	"frikiapi/src/utils/consts"
	"frikiapi/src/utils/errors"
	jwtauth "frikiapi/src/utils/jwt"
)

func (u AuthUseCases) Login(code string) (string, bool, error) {
	externalToken, err := u.AuthRepository.GenerateExternalToken(code)
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
