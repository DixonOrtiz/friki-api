package addressrepo

import (
	"gorm.io/gorm"
)

func MakeAddressRepository(DB *gorm.DB) AddressRepositoryInterface {
	return AddressRepository{
		DB: DB,
	}
}
