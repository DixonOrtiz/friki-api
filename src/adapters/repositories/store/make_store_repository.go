package storerepo

import (
	"gorm.io/gorm"
)

func MakeStoreRepository(DB *gorm.DB) IStoreRepository {
	return StoreRepository{
		DB: DB,
	}
}
