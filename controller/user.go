package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"hellonil/dao/mysql"
	"hellonil/models"
	"hellonil/pkg/jwt"
	"hellonil/responseStruct"
	"net/http"
	"strconv"
)

func responseLogin(c *gin.Context, codeErr int, userId int64, token string, statusCode int) {
	zap.L().Info(codeString[codeErr])
	c.JSON(http.StatusOK, gin.H{
		"status_code": statusCode,
		"status_msg":  codeString[codeErr],
		"user_id":     userId,
		"token":       token,
	})
}

// 注册
func Register(c *gin.Context) {
	//获取用户名或密码
	username := c.Query("username")
	password := c.Query("password")
	//判断长度
	if len(username) >= 32 || len(password) >= 32 {
		responseLogin(c, CodeUserToLength, 0, "", CodeStatusFail)
		return
	}
	accounts := &models.Accounts{ //新建一个account用户结构体
		Username: username,
		Password: password,
	}
	if !mysql.CheckUserExist(accounts.Username) { //判断用户是否存在，不存在就新建
		err := mysql.InsertAccounts(accounts)
		if err != nil {
			return
		}
	} else {
		responseLogin(c, CodeUserExist, accounts.ID, "", CodeStatusFail)
		return
	}
	tk, err := jwt.GenToken(accounts.ID, username) //生成token
	if err != nil {
		//日志
		return
	}
	//将用户信息插入到users表中
	err = mysql.InsertUsers(accounts)
	if err != nil {
		responseLogin(c, CodeServerBusy, 0, "", CodeStatusFail) //返回服务器信息繁忙
		return
	}
	responseLogin(c, CodeRegisterOk, accounts.ID, tk, CodeStatusOK)
	c.Set(tk, username)
	return
}

// 登录
func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	if len(username) >= 32 || len(password) >= 32 {
		responseLogin(c, CodeUserToLength, 0, "", CodeStatusFail)
		return
	}
	accounts := &models.Accounts{ //新建一个account用户结构体
		Username: username,
		Password: password,
	}
	pd := password
	if !mysql.CheckUserExist(accounts.Username) { //如果不存在返回错误
		responseLogin(c, CodeUserNotExist, 0, "", CodeStatusFail)
	}
	accounts.Password = pd
	//登录
	err := mysql.CheckPassWord(accounts)
	if err != nil {
		responseLogin(c, CodeInvalidPassword, 0, "", CodeStatusFail)
		return
	}
	tk, err := jwt.GenToken(accounts.ID, username)
	if err != nil {
		zap.L().Fatal("用户登录生成token失败")
		return
	}
	responseLogin(c, CodeLoginOk, accounts.ID, tk, CodeStatusOK)
}

func responsePublish(c *gin.Context, statusCode int32, statusMsg string, videoList []*responseStruct.Video) {
	c.JSON(http.StatusOK, gin.H{
		"status_code": statusCode,
		"status_msg":  statusMsg,
		"video_list":  videoList,
	})
}

// 发布列表
func PublishList(c *gin.Context) {
	token, user_id := c.Query("token"), c.Query("user_id")
	uid, err := strconv.Atoi(user_id)
	if err != nil {
		zap.L().Info("user_id含有非法字符！")
		responsePublish(c, 1, "参数错误！请重新请求", nil)
		return
	}
	myc, err := jwt.ParseToken(token)
	if err != nil {
		responsePublish(c, 1, "用户信息已过期。请重新登录！", nil)
		return
	}
	//查看用户名和user_id是否匹配
	if !mysql.UnIsOkUID(myc.Username, int64(uid)) {
		responsePublish(c, 1, "用户信息有误，请重新请求！", nil)
		return
	}
	if err != nil {
		return
	}
	vlist, err := mysql.PublishList(int64(uid))
	if err != nil {
		responsePublish(c, 1, "用户信息有误，请重新请求！", nil)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "查询成功！",
		"video_list":  vlist,
	})
}
func responseUserMsg(c *gin.Context, statusCode int32, statusMsg string, user *responseStruct.User) {
	c.JSON(http.StatusOK, gin.H{
		"status_code": statusCode,
		"status_msg":  statusMsg,
		"user":        user,
	})
}

func UserMsg(c *gin.Context) {
	token, user_id := c.Query("token"), c.Query("user_id")
	uid, err := strconv.Atoi(user_id)
	if err != nil {
		zap.L().Info("user_id含有非法字符！")
		responseUserMsg(c, 1, "参数错误！请重新请求", nil)
		return
	}
	myc, err := jwt.ParseToken(token)
	if err != nil {
		responseUserMsg(c, 1, "用户信息已过期。请重新登录！", nil)
		return
	}
	//查看用户名和user_id是否匹配
	if !mysql.UnIsOkUID(myc.Username, int64(uid)) {
		responseUserMsg(c, 1, "用户信息有误，请重新请求！", nil)
		return
	}
	fmt.Println(uid)
	res, err := mysql.SearchUserMsg(uid)
	if err != nil {
		responseUserMsg(c, 1, "用户信息有误", nil)
		return
	}
	responseUserMsg(c, 0, "成功！", res)
}
