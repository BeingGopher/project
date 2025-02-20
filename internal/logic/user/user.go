package user

import (
	"fmt"
	"project/internal/dao"
	"project/internal/model/entity"
	"project/internal/service"
)

func Register(username, password, email string) error {
	user := entity.User{Username: username, Password: password, Email: email}
	return dao.CreateUser(user)
}

func Login(username, password string) (string, error) {
	var (
		user *entity.User
		err  error
	)

	user, err = dao.GetUserByUsername(username)
	if err != nil {
		return "", fmt.Errorf("用户名错误: %v", err)
	}
	if user.Password != password {
		return "", fmt.Errorf("密码错误")
	}
	return service.GenerateToken(user)
}
