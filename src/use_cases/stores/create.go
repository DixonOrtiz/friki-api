package storeusecases

import (
	"fmt"
	"frikiapi/src/entities"
	"frikiapi/src/utils/errors"
	"frikiapi/src/utils/permissions"
	"time"

	"frikiapi/src/utils/uuid"
)

func (u StoreUseCases) Create(store entities.Store) (entities.Store, error) {
	foundStore, _, err := u.StoreRepository.GetByID(store.ID)
	if err != nil {
		return entities.Store{}, errors.New(errors.INTERNAL, err)
	}

	if foundStore.ID != "" {
		return entities.Store{}, fmt.Errorf(
			"%s: store with id %s already exists", errors.NOT_FOUND, store.ID,
		)
	}

	now := time.Now()

	store.ID = uuid.New()
	store.UpdatedAt = now
	store.CreatedAt = now

	err = u.StoreRepository.Create(store)
	if err != nil {
		return entities.Store{}, errors.New(errors.INTERNAL, err)
	}

	err = u.PermissionUseCases.AddResource(
		permissions.STORE,
		store.UserID,
		store.ID,
	)
	if err != nil {
		return entities.Store{}, errors.New(errors.INTERNAL, err)
	}

	return store, nil
}
