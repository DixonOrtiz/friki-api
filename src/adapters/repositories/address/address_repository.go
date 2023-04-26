package addressrepo

import "gorm.io/gorm"

type AddressRepository struct {
	DB *gorm.DB
}
