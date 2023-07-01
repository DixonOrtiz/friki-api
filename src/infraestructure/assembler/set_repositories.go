package assembler

import (
	oauthrepo "frikiapi/src/adapters/repositories/oauth"
)

func (a *Assembler) setRepositories() {
	// a.setUserRepository()
	a.setOAuthRepository()
	// a.setAddressRepository()
	// a.setStoreRepository()
}

func (a *Assembler) setOAuthRepository() {
	a.repositories.OAuth = oauthrepo.MakeOAuthRepository(
		a.infraestructure.OAuthConfig,
	)
}

// func (a *Assembler) setUserRepository() {
// 	a.repositories.User = userrepo.MakeUserRepository(
// 	)
// }

// func (a *Assembler) setAddressRepository() {
// 	a.repositories.Address = addressrepo.MakeAddressRepository(
// 	)
// }

// func (a *Assembler) setStoreRepository() {
// 	a.repositories.Store = storerepo.MakeStoreRepository(
// 	)
// }
