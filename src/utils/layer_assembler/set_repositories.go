package assembler

import (
	authrepo "frikiapi/src/adapters/repositories/auth"
	userrepo "frikiapi/src/adapters/repositories/user"
)

func (a *LayerAssembler) setRepositories() {
	a.setUserRepository()
	a.setAuthRepository()
}

func (a *LayerAssembler) setAuthRepository() {
	a.repositories.Auth = authrepo.MakeAuthRepository(
		a.infraestructure.AuthConfig,
	)
}

func (a *LayerAssembler) setUserRepository() {
	a.repositories.User = userrepo.MakeUserRepository(
		a.infraestructure.DB,
	)
}
