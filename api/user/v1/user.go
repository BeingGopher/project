package v1

import (
	"github.com/gin-gonic/gin"
	"project/internal/controller/user"
)

func SetUserRoute(r *gin.Engine) {
	var userGroup *gin.RouterGroup

	//r.Use(func(c *gin.Context) {
	//	c.Set("db", dao.DB)
	//	c.Next()
	//})
	userGroup = r.Group("/user")
	{
		userGroup.POST("/register", user.RegisterHandler)
		userGroup.POST("/login", user.LoginHandler)
		//userGroup.GET("/register", user.RegisterHandler)
	}
}
