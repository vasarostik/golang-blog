package gorm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func New() (*gorm.DB, error) {

	db, err := gorm.Open("postgres", "postgres://postgres:l@localhost:5432/postgres")


	if err != nil {
		return nil, err
	}

	return db, nil
}
