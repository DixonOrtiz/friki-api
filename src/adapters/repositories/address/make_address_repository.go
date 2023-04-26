package addressrepo

import (
	"gorm.io/gorm"
)

func MakeAddressRepository(DB *gorm.DB) AddressRepository {
	return AddressRepository{
		DB: DB,
	}
}
