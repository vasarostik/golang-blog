package gorm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func New(connString string) (*gorm.DB, error) {


	db, err := gorm.Open("postgres", connString)


	if err != nil {
		return nil, err
	}

	return db, nil
}
