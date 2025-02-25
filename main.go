package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	API "project/api/user/v1"
	"project/hack/config"
	"project/internal/consts"
	"project/internal/dao"
	_ "project/internal/logic"
	"project/internal/model/entity"
)

func main() {
	config.InitConfig(consts.ConfigPath)

	if err := dao.InitDB(); err != nil {
		log.Fatalf("Failed to init DB %v", err)
	}
	err := dao.DB(context.Background()).AutoMigrate(&entity.User{})
	if err != nil {
		return
	}
	r := gin.Default()
	API.SetUserRoute(r) //注册api下的路由
	errs := r.Run(":8080")
	if errs != nil {
		return
	}
	//cmd.StartServer()
}
