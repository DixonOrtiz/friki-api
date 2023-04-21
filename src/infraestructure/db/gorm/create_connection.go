package gormdb

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateConnection() (*gorm.DB, error) {
	conf, err := GetConfig()
	if err != nil {
		return nil, err
	}

	source := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.DBName,
	)

	DB, err := gorm.Open(mysql.Open(source), &gorm.Config{
		CreateBatchSize: 1000,
	})
	if err != nil {
		return nil, err
	}

	return DB, nil
}
