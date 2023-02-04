package middlewares

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"hellonil/pkg/jwt"
	"net/http"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Query("token")
		if len(auth) == 0 {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"status_code": 0,
				"status_msg":  "请重新登录！",
			})
			return
		}
		// 校验token，只要出错直接拒绝请求
		myC, err := jwt.ParseToken(auth)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"status_code": 0,
				"status_msg":  "请重新登录！",
			})
			return
		}
		zap.L().Info("用户鉴权成功,用户名为", zap.Stack(myC.Username))
		c.Next()
	}
}
