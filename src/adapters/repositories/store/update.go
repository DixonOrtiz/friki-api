package storerepo

import "frikiapi/src/entities"

func (r StoreRepository) Update(store entities.Store) error {
	err := r.DB.
		Table("stores").
		Omit("external_id").
		Updates(store).
		Where("id = ?", store.ID).Error
	if err != nil {
		return err
	}

	return nil
}
