package assembler

import (
	addressrepo "frikiapi/src/adapters/repositories/address"
	oauthrepo "frikiapi/src/adapters/repositories/oauth"
	storerepo "frikiapi/src/adapters/repositories/store"
	userrepo "frikiapi/src/adapters/repositories/user"
)

func (a *LayerAssembler) setRepositories() {
	a.setUserRepository()
	a.setOAuthRepository()
	a.setAddressRepository()
	a.setStoreRepository()
}

func (a *LayerAssembler) setOAuthRepository() {
	a.repositories.OAuth = oauthrepo.MakeOAuthRepository(
		a.infraestructure.OAuthConfig,
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
