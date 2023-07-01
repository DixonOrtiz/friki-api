package assembler

import (
	oauthrepo "frikiapi/src/adapters/repositories/oauth"
	googleuserrepo "frikiapi/src/adapters/repositories/user/google"
)

func (a *Assembler) setRepositories() {
	// a.setUserRepository()
	a.setOAuthRepository()
	a.setExternalUserRepository()
	// a.setAddressRepository()
	// a.setStoreRepository()
}

func (a *Assembler) setOAuthRepository() {
	a.repositories.OAuth = oauthrepo.MakeOAuthRepository(
		a.infraestructure.OAuthConfig,
	)
}

func (a *Assembler) setExternalUserRepository() {
	a.repositories.ExternalUser = googleuserrepo.MakeUserRepository()
}

// func (a *Assembler) setAddressRepository() {
// 	a.repositories.Address = addressrepo.MakeAddressRepository(
// 	)
// }

// func (a *Assembler) setStoreRepository() {
// 	a.repositories.Store = storerepo.MakeStoreRepository(
// 	)
// }
