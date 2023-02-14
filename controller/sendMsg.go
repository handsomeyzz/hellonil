package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"hellonil/dao/mysql"
	"net/http"
)

func responseSendMsg(c *gin.Context, StatusCode int, StatusMsg string) {
	c.JSON(http.StatusOK, gin.H{
		"status_code": StatusCode,
		"status_msg":  StatusMsg,
	})
}

func SendMsg(c *gin.Context) {
	token, to_user_id, action_type, content := c.Query("token"), c.Query("to_user_id"), c.Query("action_type"), c.Query("content")
	from_user_id, exist := c.Get(token)
	if !exist {
		zap.L().Info("token有误")
		responseSendMsg(c, 1, "发送消息失败")
		return
	}
	if action_type != "1" {
		zap.L().Info("参数传入错误")
		responseSendMsg(c, 1, "发送消息失败")
		return
	}
	err := mysql.InsertMsg(from_user_id, to_user_id, content)
	if err != nil {
		responseSendMsg(c, 1, "发送消息失败")
		return
	}
	responseSendMsg(c, 0, "发送成功!")
}
