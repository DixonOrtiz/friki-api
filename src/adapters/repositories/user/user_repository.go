package userrepo

import "cloud.google.com/go/firestore"

type UserRepository struct {
	DB *firestore.Client
}
