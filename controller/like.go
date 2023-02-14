package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"hellonil/dao/mysql"
	"hellonil/pkg/jwt"
	"hellonil/responseStruct"
	"net/http"
)

func response(c *gin.Context, StatusCode string, StatusMsg string, VideoList []*responseStruct.Video) {
	c.JSON(http.StatusOK, gin.H{
		"status_code": StatusCode,
		"status_msg":  StatusMsg,
		"video_list":  VideoList,
	})
}

func FavoriteList(c *gin.Context) {
	user_id, token := c.Query("user_id"), c.Query("token")
	_, err := jwt.ParseToken(token)
	if err != nil {
		zap.L().Info("token有错误，请重新操作")
		response(c, "1", "请重新尝试", nil)
		return
	}
	//if strconv.Itoa(int(myc.UserID)) != user_id {
	//	zap.L().Info("token中的UserID与传入的user_id有误")
	//	response(c, "1", "请重新尝试", nil)
	//	return
	//}

	res, err := mysql.LikeList(user_id)
	if err != nil {
		zap.L().Info("查询错误")
		response(c, "1", "请重新尝试", nil)
		return
	}
	response(c, "0", "查询成功", res)
	return
}
