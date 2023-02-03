package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"hellonil/dao/mysql"
	"hellonil/models"
	"hellonil/pkg/jwt"
	"net/http"
)

func responseLogin(c *gin.Context, codeErr int, userId int64, token string) {
	zap.L().Info(codeString[codeErr])
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  codeString[codeErr],
		"user_id":     userId,
		"token":       token,
	})
}

func Register(c *gin.Context) {
	//获取用户名或密码
	username := c.Query("username")
	password := c.Query("password")
	//判断长度
	if len(username) >= 32 || len(password) >= 32 {
		responseLogin(c, CodeUserToLength, 0, "")
		return
	}
	accounts := &models.Accounts{ //新建一个account用户结构体
		Username: username,
		Password: password,
	}
	if !mysql.CheckUserExist(accounts) { //判断用户是否存在，不存在就新建
		err := mysql.InsertAccounts(accounts)
		if err != nil {
			return
		}
	} else {
		responseLogin(c, CodeUserExist, accounts.ID, "")
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
		responseLogin(c, CodeServerBusy, 0, "") //返回服务器信息繁忙
		return
	}
	responseLogin(c, CodeRegisterOk, accounts.ID, tk)
	c.Set(tk, username)
	return
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	if len(username) >= 32 || len(password) >= 32 {
		responseLogin(c, CodeUserToLength, 0, "")
		return
	}
	accounts := &models.Accounts{ //新建一个account用户结构体
		Username: username,
		Password: password,
	}
	pd := password
	if !mysql.CheckUserExist(accounts) { //如果不存在返回错误
		responseLogin(c, CodeUserNotExist, 0, "")
	}
	accounts.Password = pd
	//登录
	err := mysql.CheckPassWord(accounts)
	if err != nil {
		responseLogin(c, CodeInvalidPassword, 0, "")
		return
	}
	tk, err := jwt.GenToken(accounts.ID, username)
	if err != nil {
		zap.L().Fatal("用户登录生成token失败")
		return
	}
	responseLogin(c, CodeLoginOk, accounts.ID, tk)
}
