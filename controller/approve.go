package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"hellonil/dao/mysql"
	"hellonil/pkg/jwt"
	"net/http"
)

func Approve(c *gin.Context) {
	token := c.Query("token")      //人
	videoId := c.Query("video_id") //视频id
	actionType := c.Query("action_type")
	//根据token获取userid
	myc, err := jwt.ParseToken(token)
	StatusCode := 0
	str := "服务器出错"
	if err != nil {
		StatusCode = 1
		zap.L().Info("解析token出错！")
	}
	userId := myc.UserID
	//对数据库进行操作
	//查询该视频点赞数 n为点赞数
	n, err := mysql.InquireFavorite(videoId)
	if err != nil {
		StatusCode = 2
		zap.L().Info("inquire favorite failed")
		return
	}
	if actionType == "1" {
		//插入喜欢
		err := mysql.InsertLike(userId, videoId)
		if err != nil {
			StatusCode = 3
			zap.L().Info("insert like.go failed")
		}
		//将视频点赞数加1
		err2 := mysql.UpdateFavorite(videoId, n+1)
		if err2 != nil {
			StatusCode = 4
			zap.L().Info("update favoriteCount failed")
		}
		str = "点赞成功"
	} else {
		//删除喜欢记录
		err := mysql.DeleteLike(userId, videoId)
		if err != nil {
			StatusCode = 5
			zap.L().Info("delete like.go failed")
		}
		//取消赞
		err2 := mysql.UpdateFavorite(videoId, n-1)
		if err2 != nil {
			StatusCode = 6
			zap.L().Info("cancel favorite failed")
		}
		str = "取消成功"
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code": StatusCode,
		"status_msg":  str,
	})
}
