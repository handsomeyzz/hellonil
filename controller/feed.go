package controller

import (
	"github.com/gin-gonic/gin"
	"hellonil/dao/mysql"
	"time"
)

type Feed struct {
	StatusCode int           `json:"status_code"`
	StatusMsg  string        `json:"status_msg"`
	NextTime   time.Time     `json:"next_time"`
	VideoList  []mysql.Video `json:"video_list"`
}

func FeedSearch(c *gin.Context) {
	//latest_time := c.Query("latest_time")
	token := c.Query("token")
	myfeed := make([]*mysql.Video, 0, 30)
	//fmt.Println("token:", token, "latest_time:", latest_time)
	isOk := true
	if token == "" {
		isOk = false
	}
	err := mysql.SearchFeed(myfeed, isOk)
	if err != nil {
		return
	}

}
