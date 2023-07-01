package addressrepo

import (
	"gorm.io/gorm"
)

func MakeAddressRepository(DB *gorm.DB) IAddressRepository {
	return AddressRepository{
		DB: DB,
	}
}
