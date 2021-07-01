package middlewares

import (
	"blog-backend/pkg/jwt"
	"blog-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从 header 中获取 token
		token := c.Request.Header.Get("token")

		if token == "" {
			c.Abort()
			response.UnAuthorized(c, "header缺少token")
			return
		}

		claims, err := jwt.ParseToken(token)

		if err != nil {
			c.Abort()
			response.UnAuthorized(c, err.Error())
			return
		}

		c.Set("UserId", claims.ID)

		c.Next()
	}
}
