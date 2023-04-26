package storerepo

import "frikiapi/src/entities"

func (r StoreRepository) Create(store entities.Store) (int, error) {
	err := r.DB.Table("stores").Create(&store).Error
	if err != nil {
		return 0, err
	}

	return store.ID, nil
}
