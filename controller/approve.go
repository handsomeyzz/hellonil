package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Approve(c *gin.Context) {
	token := c.Query("token")       //人
	video_id := c.Query("video_id") //视频id
	action_type := c.Query("action_type")
	fmt.Println(token, video_id, action_type)
	//对数据库进行操作
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "点赞成功",
	})

}
