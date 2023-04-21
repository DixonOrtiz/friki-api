package authrepo

import authinfra "frikiapi/src/infraestructure/auth"

type AuthRepository struct {
	Config authinfra.OAuth2ConfigInterface
}
