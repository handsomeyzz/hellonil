package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"hellonil/dao/mysql"
	"hellonil/pkg/jwt"
	"hellonil/responseStruct"
	"net/http"
	"strconv"
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
		rFollow(c, 1, "关注失败")
		return
	}
	username := myc.Username
	if !mysql.CheckUserExist(username) {
		zap.L().Info("用户不存在")
		rFollow(c, 1, "用户信息有误")
		return
	}
	if !mysql.FollowCheckUid(to_user_id) {
		zap.L().Info("用户不存在")
		rFollow(c, 1, "用户信息有误")
		return
	}
	if action_type == "1" {
		err = mysql.FollowInsert(username, to_user_id)
		if err != nil {
			rFollow(c, 1, "关注失败")
			return
		}
	} else if action_type == "2" {
		err = mysql.FollowDelete(username, to_user_id)
		if err != nil {
			rFollow(c, 1, "关注失败")
			return
		}
	} else {
		rFollow(c, 1, "关注失败")
		return
	}
	rFollow(c, 0, "关注成功")
	return
}

func responseFollowList(c *gin.Context, StatusCode int, StatusMsg string, UserList []*responseStruct.User) {
	c.JSON(http.StatusOK, gin.H{
		"status_code": StatusCode,
		"status_msg":  StatusMsg,
		"user_list":   UserList,
	})
}

// 粉丝列表
func FollowList(c *gin.Context) {
	user_id, token := c.Query("user_id"), c.Query("token")
	myc, err := jwt.ParseToken(token)
	if err != nil {
		zap.L().Info("解析token出错", zap.Error(err))
		responseFollowList(c, 1, "用户信息有误", nil)
		return
	}
	if user_id != strconv.Itoa(int(myc.UserID)) {
		zap.L().Info("token和用户名不匹配")
		responseFollowList(c, 1, "用户信息有误", nil)
		return
	}
	res, err := mysql.FollowList(user_id)
	if err != nil {
		responseFollowList(c, 1, "用户信息有误", nil)
		return
	}
	responseFollowList(c, 0, "OK", res)
	return
}
