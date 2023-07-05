package assembler

import (
	oauthrepo "frikiapi/src/adapters/repositories/oauth"
	userrepo "frikiapi/src/adapters/repositories/user"
	googleuserrepo "frikiapi/src/adapters/repositories/user/google"
)

func (a *Assembler) setRepositories() {
	a.setUserRepository()
	a.setOAuthRepository()
	a.setExternalUserRepository()
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
