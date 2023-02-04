package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"hellonil/dao/feed"
	"hellonil/dao/mysql"
	"hellonil/models"
	"hellonil/pkg/jwt"
	"net/http"
	"os"
	"time"
)

func publish(c *gin.Context, msg string, statusCode int) {
	c.JSON(http.StatusOK, gin.H{
		"status_code": statusCode,
		"status_msg":  msg,
	})
}

func PublishAction(c *gin.Context) {
	//获取数据源
	form, err := c.MultipartForm()
	if err != nil {
		return
	}
	v := form.File["data"]
	tk := form.Value["token"]
	tt := form.Value["title"]
	if len(v) != 1 || len(tk) != 1 || len(tt) != 1 {
		zap.L().Info("上传视频失败！参数不符合规定")
		publish(c, "文件上传失败!参数不符合规定", CodeStatusFail)
		return
	}
	video := v[0]
	token := tk[0]
	title := tt[0]
	if token == "" {
		zap.L().Info("上传视频时token校验失败")
		publish(c, "文件上传失败!参数不符合规定", CodeStatusFail)
		return
	} else { //鉴权
		myC, err := jwt.ParseToken(token)
		if err != nil {
			zap.L().Info("上传视频时token校验失败")
			publish(c, "文件上传失败!参数不符合规定", CodeStatusFail)
			return
		}
		zap.L().Info("用户鉴权成功,用户名为", zap.Stack(myC.Username))
	}
	if err != nil {
		zap.L().Info("err:up video failed and ", zap.Error(err))
		publish(c, "文件上传失败!", CodeStatusFail)
		return
	}
	mc, _ := jwt.ParseToken(token)
	username := mc.Username
	dst := "./video/" + title + ".mp4"
	//保存视频
	err = c.SaveUploadedFile(video, dst)
	if err != nil {
		zap.L().Info("err:save file and ", zap.Error(err))
		publish(c, "文件上传失败!", CodeStatusFail)
		return
	}
	videoPath, picPath, err := feed.DealVideo(dst)
	if err != nil {
		zap.L().Info("文件处理失败")
		publish(c, "文件上传失败!", CodeStatusFail)
		return
	}
	//上传视频
	vname := username[:4] + title + ".mp4" //视频名字
	pname := username[:4] + title + ".jpg" //视频的图片
	videoUrl, err := feed.Upload(vname, videoPath)
	picUrl, err := feed.Upload(pname, picPath)
	err = os.Remove(dst)
	videoStruct := &models.Videos{
		AuthorID:      mc.UserID,
		PlayUrl:       videoUrl,
		CoverUrl:      picUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         title,
		CreateTime:    time.Now(),
	}
	err = mysql.InsertVideo(videoStruct)
	if err != nil {
		zap.L().Info("文件上传时数据插入失败")
		publish(c, "文件上传失败", CodeStatusFail)
		return
	}
	publish(c, "上传成功！", CodeStatusOK)
	fmt.Println("上传欧克")
}
