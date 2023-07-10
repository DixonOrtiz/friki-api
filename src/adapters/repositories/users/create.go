package userrepo

import (
	"context"
	"frikiapi/src/adapters/repositories/users/types"
	"frikiapi/src/entities"
)

func (r *UserRepository) Create(user entities.User) error {
	firestoreUser := types.MapUserToFirestore(user)

	_, err := r.DB.Collection("users").NewDoc().Create(context.Background(), firestoreUser)
	if err != nil {
		return err
	}

	return nil
}
