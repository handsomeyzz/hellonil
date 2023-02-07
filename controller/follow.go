package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"hellonil/dao/mysql"
	"hellonil/pkg/jwt"
	"net/http"
)

func rFollow(c *gin.Context, StatusCode int, StatusMsg string) {
	c.JSON(http.StatusOK, gin.H{
		"status_code": StatusCode,
		"status_msg":  StatusMsg,
	})
}

func FollowAction(c *gin.Context) {
	token, to_user_id, action_type := c.Query("token"), c.Query("to_user_id"), c.Query("action_type")
	myc, err := jwt.ParseToken(token)
	if err != nil {
		zap.L().Info("关注操作解析token失败", zap.Error(err))
		rFollow(c,1,"关注失败")
		return
	}
	username:=myc.Username
	if !mysql.CheckUserExist(username) || !mysql.CheckUserExist()

}
