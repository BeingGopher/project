package dao

import (
	"context"
	"project/internal/model/entity"
)

func CreateUser(user entity.User) error {

	return DB(context.Background()).Create(&user).Error
}

func GetUserByUsername(username string) (*entity.User, error) {
	var (
		user entity.User
		err  error
	)
	err = DB(context.Background()).Where("username = ?", username).First(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, nil
}
