package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"hellonil/dao/feed"
	"net/http"
	"os"
)

func PublishAction(c *gin.Context) {
	video, err := c.FormFile("data")
	if err != nil {
		zap.L().Fatal("err:up video failed and ", zap.Error(err))
		c.JSON(http.StatusOK, "failed!")
	}

	title := c.Request.Header.Get("title")

	dst := "./video/" + title + ".mp4"
	//保存视频
	err = c.SaveUploadedFile(video, dst)
	if err != nil {
		zap.L().Fatal("err:save file and ", zap.Error(err))
		c.JSON(http.StatusOK, "upload file failed")
	}

	videoPath, picPath, err := feed.DealVideo(dst)
	if err != nil {
		fmt.Println(err)
	}
	videoUrl, err := feed.Upload(title+".mp4", videoPath)
	picUrl, err := feed.Upload(title+".jpg", picPath)
	err = os.Remove(dst)
	if err != nil {

	}
	fmt.Println(videoUrl, picUrl)
}
