package userrepo

import (
	"context"
	"frikiapi/src/adapters/repositories/user/types"
	"frikiapi/src/entities"
	"time"
)

func (r *UserRepository) Create(user entities.User) error {
	firestoreUser := types.MapUserToFirestore(user)

	now := time.Now()
	firestoreUser.UpdatedAt = now
	firestoreUser.CreatedAt = now

	_, err := r.DB.Collection("users").NewDoc().Create(context.Background(), firestoreUser)
	if err != nil {
		return err
	}

	return nil
}
