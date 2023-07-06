package userrepo

import (
	"context"
	"frikiapi/src/adapters/repositories/user/types"
	"frikiapi/src/entities"
)

func (r *UserRepository) Update(document string, user entities.User) error {
	firestoreUser := types.MapUserToFirestore(user)

	_, err := r.DB.Collection("users").Doc(document).Set(
		context.Background(),
		firestoreUser,
	)
	if err != nil {
		return err
	}

	return nil
}
