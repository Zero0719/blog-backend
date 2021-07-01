package routers

import (
	"blog-backend/controllers"
	"blog-backend/middlewares"

	"github.com/gin-gonic/gin"
)

// @Title SetUp
// @Description 路由初始化
func SetUp() *gin.Engine {
	router := gin.Default()

	router.POST("/user", controllers.UserRegister) // 用户注册
	router.POST("/auth", controllers.UserLogin)    // 用户登录

	router.GET("/users/:id", controllers.UserShow) // 单个用户信息

	// 需要登录验证权限的路由
	auth := router.Group("")
	auth.Use(middlewares.Auth())
	auth.POST("/follow/:id", controllers.UserFollow) // 关注
	auth.POST("/unfollow/:id", controllers.UserUnFollow) // 取关
	auth.POST("/articles", controllers.ArticleStore) // 发布文章

	return router
}
