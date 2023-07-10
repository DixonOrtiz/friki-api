package googleuserrepo

func MakeUserRepository() IGoogleUserRepository {
	return &GoogleUserRepository{}
}
