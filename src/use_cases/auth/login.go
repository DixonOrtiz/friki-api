package authusecases

import (
	"frikiapi/src/entities"
	"frikiapi/src/utils/consts"
	"frikiapi/src/utils/errors"
	jwtauth "frikiapi/src/utils/jwt"
)

func (u AuthUseCases) Login(code string) (string, error) {
	tokenStr, err := u.AuthRepository.GenerateExternalToken(code)
	if err != nil {
		return "", errors.NewError(consts.INTERNAL, err)
	}

	externalUser, err := u.UserRepository.GetExternalUserByToken(tokenStr)
	if err != nil {
		return "", errors.NewError(consts.INTERNAL, err)
	}

	user, err := u.UserRepository.GetByExternalID(externalUser.ExternalID)
	if err != nil {
		return "", errors.NewError(consts.INTERNAL, err)
	}

	if user == (entities.User{}) {
		externalUser.RoleID = consts.USER
		err = u.UserRepository.Register(externalUser)
		if err != nil {
			return "", errors.NewError(consts.INTERNAL, err)
		}
	}

	token, err := jwtauth.GenerateJWT(externalUser.ExternalID)
	if err != nil {
		return "", errors.NewError(consts.INTERNAL, err)
	}

	return token, nil
}