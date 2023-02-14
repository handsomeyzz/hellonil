package middlewares

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"hellonil/pkg/jwt"
	"net/http"
)

func re(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "请重新登录！",
	})
}

// 只适用于两个返回值的场合
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Query("token")
		// 校验token，只要出错直接拒绝请求
		myC, err := jwt.ParseToken(auth)
		if err != nil {
			c.Abort()
			zap.L().Info("token解析失败！")
			re(c)
			return
		}
		zap.L().Info("用户鉴权成功,用户名为", zap.Stack(myC.Username))
		c.Set(auth, myC.UserID)
		c.Next()
	}
}
