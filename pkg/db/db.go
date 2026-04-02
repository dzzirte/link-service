package db

import (
	"awesomeproject/configs"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func NewDb(conf *configs.Config) *Db {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		connStr = conf.Db.DSN
	}
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return &Db{db}
}
