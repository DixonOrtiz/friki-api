package addressrepo

import "frikiapi/src/entities"

func (r AddressRepository) Create(address entities.Address) (int, error) {
	err := r.DB.Table("addresses").Create(&address).Error
	if err != nil {
		return 0, err
	}

	return address.ID, nil
}
