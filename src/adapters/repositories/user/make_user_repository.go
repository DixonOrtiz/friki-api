package userrepo

import "cloud.google.com/go/firestore"

func MakeUserRepository(DB *firestore.Client) IUserRepository {
	return &UserRepository{
		DB: DB,
	}
}
