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

	router.GET("/users/:id", controllers.UserShow)                // 单个用户信息
	router.GET("/articles/:id", controllers.ArticleShow)          // 单片文章信息
	router.GET("/articles", controllers.ArticleIndex)             // 文章列表
	router.GET("/comments/:article_id", controllers.CommentIndex) // 评论列表

	// 需要登录验证权限的路由
	auth := router.Group("")
	auth.Use(middlewares.Auth())
	auth.POST("/follow/:id", controllers.UserFollow)             // 关注
	auth.POST("/unfollow/:id", controllers.UserUnFollow)         // 取关
	auth.POST("/articles", controllers.ArticleStore)             // 发布文章
	auth.PUT("/articles/:id", controllers.ArticleUpdate)         // 修改文章
	auth.DELETE("/articles/:id", controllers.ArticleDestroy)     // 文章删除
	auth.POST("/comments/:article_id", controllers.CommentStore) // 发表评论
	auth.DELETE("/comments/:id", controllers.CommentDestroy)     // 删除评论

	return router
}
