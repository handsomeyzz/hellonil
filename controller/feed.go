package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"hellonil/dao/mysql"
	"hellonil/responseStruct"
	"net/http"
)

func FeedSearch(c *gin.Context) {
	//latest_time := c.Query("latest_time")
	token := c.Query("token")
	isOk := true
	if token == "" {
		isOk = false
	}
	vStruct, err := mysql.SearchFeed(isOk, token)
	if err != nil {
		return
	}
	res := responseStruct.DouYinResponse{
		StatusCode: CodeStatusOK,
		StatusMsg:  "加载成功",
		VideoList:  vStruct,
		NextTime:   100000,
	}
	//re, err := json.Marshal(res)
	if err != nil {
		zap.L().Info("json转化失败")
		return
	}
	c.JSON(http.StatusOK, res)
}
