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

func reFollower(c *gin.Context, StatusCode, StatusMsg string, UserList []*responseStruct.User) {
	c.JSON(http.StatusOK, gin.H{
		"status_code": StatusCode,
		"status_msg":  StatusMsg,
		"user_list":   UserList,
	})
}

func FollowerList(c *gin.Context) {
	token, user_id := c.Query("token"), c.Query("user_id")
	myc, err := jwt.ParseToken(token)
	if err != nil || user_id != strconv.Itoa(int(myc.UserID)) {
		zap.L().Info("token解析失败!")
		reFollower(c, "1", "用户信息有误", nil)
		return
	}
	res, err := mysql.FollowerList(user_id)
	if err != nil {
		reFollower(c, "1", "用户信息有误", nil)
		return
	}
	reFollower(c, "0", "OK", res)
}
