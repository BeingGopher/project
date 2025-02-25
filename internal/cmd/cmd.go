package cmd

import (
	"github.com/gin-gonic/gin"
	"log"
	API "project/api/user/v1"
	"project/hack/config"
	"project/internal/consts"
	"project/internal/dao"
)

// InitApp 初始化应用并返回 Gin 路由
func InitApp() *gin.Engine {
	// 初始化配置

	config.InitConfig(consts.ConfigPath)

	// 初始化数据库
	if err := dao.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 数据库迁移
	//if err := dao.DB().AutoMigrate(&entity.User{}).Error; err != nil {
	//	log.Fatalf("Failed to migrate database: %v", err)
	//}

	// 初始化 Gin 路由
	r := gin.Default()
	API.SetUserRoute(r) // 注册路由

	return r
}

// StartServer 启动服务器
func StartServer() {
	router := InitApp() // 初始化应用并获取路由

	port := ":8080" // 可以从配置文件中读取端口
	router.Run(port)
	log.Printf("Starting server on %s", port)
	if err := router.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
