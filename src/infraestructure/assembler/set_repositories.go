package assembler

import (
	addressrepo "frikiapi/src/adapters/repositories/addresses"
	oauthrepo "frikiapi/src/adapters/repositories/oauth"
	permrepo "frikiapi/src/adapters/repositories/permissions"
	storerepo "frikiapi/src/adapters/repositories/stores"
	userrepo "frikiapi/src/adapters/repositories/users"
	googleuserrepo "frikiapi/src/adapters/repositories/users/google"
)

func (a *Assembler) setRepositories() {
	a.setUserRepository()
	a.setOAuthRepository()
	a.setExternalUserRepository()
	a.setAddressRepository()
	a.setPermissionRepository()
	a.setStoreRepository()
}

func (a *Assembler) setUserRepository() {
	a.repositories.User = userrepo.MakeUserRepository(a.infraestructure.DB)
}

func (a *Assembler) setOAuthRepository() {
	a.repositories.OAuth = oauthrepo.MakeOAuthRepository(
		a.infraestructure.OAuthConfig,
	)
}

func (a *Assembler) setExternalUserRepository() {
	a.repositories.ExternalUser = googleuserrepo.MakeUserRepository()
}

func (a *Assembler) setAddressRepository() {
	a.repositories.Address = addressrepo.MakeAddressRepository(a.infraestructure.DB)
}

func (a *Assembler) setPermissionRepository() {
	a.repositories.Permission = permrepo.MakePermissionRepository(a.infraestructure.DB)
}

func (a *Assembler) setStoreRepository() {
	a.repositories.Store = storerepo.MakeStoreRepository(a.infraestructure.DB)
}
