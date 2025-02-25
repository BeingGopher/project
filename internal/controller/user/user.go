package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project/internal/model/entity"
	"project/internal/service"
)

func RegisterHandler(c *gin.Context) {
	var (
		req entity.User
		err error
	)
	if err = c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if err = service.User().Register(req.Username, req.Password, req.Email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "注册成功"})

}

func LoginHandler(c *gin.Context) {
	var (
		req   entity.User
		err   error
		token string
	)
	if err = c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err = service.User().Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "无效token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
