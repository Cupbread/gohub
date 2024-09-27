package bootstrap

import (
	"errors"
	"fmt"
	"gohub/pkg/config"
	"gohub/pkg/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupDB() {
	var dbConfig gorm.Dialector
	if config.Get("database.connection") == "mysql" {
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
			config.Get("database.user"),
			config.Get("database.password"),
			config.Get("database.host"),
			config.Get("database.port"),
			config.Get("database.name"),
			config.Get("database.charset"),
		)
		dbConfig = mysql.New(mysql.Config{
			DSN: dsn,
		})
	} else {
		panic(errors.New("database connection not set"))
	}

	database.Connect(dbConfig, logger.Default.LogMode(logger.Info))

}
