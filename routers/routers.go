package routers

import (
	"blog-backend/controllers"

	"github.com/gin-gonic/gin"
)

// @Title SetUp
// @Description 路由初始化
func SetUp() *gin.Engine {
	router := gin.Default()

	router.POST("/user", controllers.UserRegister) // 用户注册
	router.POST("/auth", controllers.UserLogin)    // 用户登录

	router.GET("/users/:id", controllers.UserShow) // 单个用户信息
	//router.POST("/follow/:id", controllers.UserFollow) // 关注

	return router
}
