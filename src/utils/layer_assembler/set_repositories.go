package assembler

import (
	addressrepo "frikiapi/src/adapters/repositories/address"
	authrepo "frikiapi/src/adapters/repositories/auth"
	storerepo "frikiapi/src/adapters/repositories/store"
	userrepo "frikiapi/src/adapters/repositories/user"
)

func (a *LayerAssembler) setRepositories() {
	a.setUserRepository()
	a.setAuthRepository()
	a.setAddressRepository()
	a.setStoreRepository()
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

func (a *LayerAssembler) setAddressRepository() {
	a.repositories.Address = addressrepo.MakeAddressRepository(
		a.infraestructure.DB,
	)
}

func (a *LayerAssembler) setStoreRepository() {
	a.repositories.Store = storerepo.MakeStoreRepository(
		a.infraestructure.DB,
	)
}
