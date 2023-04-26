package storerepo

import (
	"gorm.io/gorm"
)

func MakeStoreRepository(DB *gorm.DB) StoreRepositoryInterface {
	return StoreRepository{
		DB: DB,
	}
}
