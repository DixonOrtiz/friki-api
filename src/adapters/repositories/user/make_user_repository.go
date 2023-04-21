package userrepo

import (
	"gorm.io/gorm"
)

func MakeUserRepository(DB *gorm.DB) UserRepositoryInterface {
	return &UserRepository{
		DB: DB,
	}
}
