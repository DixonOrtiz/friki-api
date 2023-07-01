package gormdb

import (
	"frikiapi/src/utils/types"
	"os"
)

var (
	DB_USER     = "DB_USER"
	DB_PASSWORD = "DB_PASSWORD"
	DB_HOST     = "DB_HOST"
	DB_NAME     = "DB_NAME"
	DB_PORT     = "DB_PORT"
)

func GetConfig() (Config, error) {
	port, err := types.ParseStringToInt(os.Getenv(DB_PORT))
	if err != nil {
		return Config{}, err
	}

	return Config{
		User:     os.Getenv(DB_USER),
		Password: os.Getenv(DB_PASSWORD),
		Host:     os.Getenv(DB_HOST),
		Port:     port,
		DBName:   os.Getenv(DB_NAME),
	}, nil
}
