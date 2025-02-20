package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"project/hack/config"
)

var DB *gorm.DB

func InitDB() (errs error) {
	var (
		dsn string
		db  *gorm.DB
		err error
	)
	dsn = config.Config.Database.DSN
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return
	}
	DB = db
	return
}
