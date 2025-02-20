package dao

import (
	"errors"
	"project/internal/model/entity"
)

func CreateUser(user entity.User) error {
	if DB == nil {
		return errors.New("数据库未连接")
	}
	return DB.Create(&user).Error
}

func GetUserByUsername(username string) (*entity.User, error) {
	var (
		user entity.User
		err  error
	)
	err = DB.Where("username = ?", username).First(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, nil
}
