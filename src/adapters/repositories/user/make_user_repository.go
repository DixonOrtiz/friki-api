package userrepo

import (
	"gorm.io/gorm"
)

func MakeUserRepository(DB *gorm.DB) IUserRepository {
	return &UserRepository{
		DB: DB,
	}
}
