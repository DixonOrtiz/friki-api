package storerepo

import "gorm.io/gorm"

type StoreRepository struct {
	DB *gorm.DB
}
