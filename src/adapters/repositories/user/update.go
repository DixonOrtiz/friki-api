package userrepo

import (
	"context"
	"frikiapi/src/adapters/repositories/user/types"
	"frikiapi/src/entities"
	"time"
)

func (r *UserRepository) Update(document string, user entities.User) error {
	firestoreUser := types.MapUserToFirestore(user)
	firestoreUser.UpdatedAt = time.Now()

	_, err := r.DB.Collection("users").Doc(document).Set(
		context.Background(),
		firestoreUser,
	)
	if err != nil {
		return err
	}

	return nil
}
