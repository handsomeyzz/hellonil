package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hellonil/dao/mysql"
	"hellonil/pkg/jwt"
	"hellonil/pkg/snowflake"
	"net/http"
)

func Register(c *gin.Context) {
	//获取用户名或密码
	username := c.Query("username")
	password := c.Query("password")
	//判断长度
	if len(username) >= 32 || len(password) >= 32 {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  CodeUserToLength,
			"user_id":     0,
			"token":       "",
		})
		return
	}

	//判断用户是否存在
	err := mysql.CheckUserExist(username)

	//雪花算法生成用户ID
	user_id := snowflake.GenID()
	//md5算法对密码加密
	tk, err := jwt.GenToken(user_id, username)
	if err != nil {
		//日志
		return
	}

	//插入数据进数据库
	err = mysql.InsertUser()

	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "注册成功",
		"user_id":     user_id,
		"token":       tk,
	})
	c.Set(tk, username)
	un, err := jwt.ParseToken(tk)
	if err != nil {
		return
	}
	fmt.Println(c.Keys[tk], un.Username)
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	if len(username) >= 32 || len(password) >= 32 {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  CodeUserToLength,
			"user_id":     0,
			"token":       "",
		})
		return
	}
}
