package storerepo

import "frikiapi/src/entities"

func (r StoreRepository) GetByExternalID(externalID string) (entities.Store, error) {
	var store entities.Store
	emptyStore := entities.Store{}

	err := r.DB.Table("stores").Where("external_id = ?", externalID).First(&store).Error

	if err != nil {
		if err.Error() == "record not found" {
			return emptyStore, nil
		}
		return emptyStore, err
	}

	if store.ID != 0 {
		return store, nil
	}

	return emptyStore, nil
}
