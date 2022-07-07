package db

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() error {
	t_db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	db = t_db

	return nil
}

func DB() *gorm.DB {
	return db
}
