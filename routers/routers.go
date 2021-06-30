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

	return router
}
