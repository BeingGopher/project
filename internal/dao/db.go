package dao

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"project/hack/config"
)

var db *gorm.DB

func InitDB() (errs error) {
	var (
		dsn string
		err error
	)
	dsn = config.Config.Database.DSN
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return
	}

	return
}

func DB(ctx context.Context) *gorm.DB {
	return db.WithContext(ctx)
}
