package pkt

import (
	"github.com/golang-jwt/jwt/v4"
	"project/internal/consts"
	"project/internal/model/entity"
	"time"
)

func GenerateToken(user *entity.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), //token有效期
	})
	return token.SignedString([]byte(consts.JwtSecret))
}
