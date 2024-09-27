package database

import (
	"fmt"
	"gorm.io/gorm"

	gormlogger "gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect(dbConfig gorm.Dialector, _logger gormlogger.Interface) {
	var err error
	DB, err = gorm.Open(dbConfig, &gorm.Config{
		Logger: _logger,
	})

	if err != nil {
		fmt.Println(err)
	}
}
