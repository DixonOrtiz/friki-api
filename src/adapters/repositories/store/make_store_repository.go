package storerepo

import (
	"gorm.io/gorm"
)

func MakeStoreRepository(DB *gorm.DB) StoreRepository {
	return StoreRepository{
		DB: DB,
	}
}
