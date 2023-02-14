package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"hellonil/dao/mysql"
	"hellonil/pkg/jwt"
	"hellonil/responseStruct"
	"net/http"
	"strconv"
)

func responseFail(c *gin.Context, statusMsg int, msg string, rsc *responseStruct.Comment) {
	c.JSON(http.StatusOK, gin.H{
		"status_code": statusMsg,
		"status_msg":  msg,
		"comment":     rsc,
	})
}

func CommentAction(c *gin.Context) {
	//1.token鉴权(判断用户存在与否)
	myc, err := jwt.ParseToken(c.Query("token")) //解析token，失败就返回并记录错误
	if err != nil {
		zap.L().Info("token有误！错误为:", zap.Error(err))
		responseFail(c, 1, "获取信息失败，请稍后重试", nil)
		return
	}

	username := myc.Username
	if !mysql.CheckUserExist(username) { //查询数据库用户存在与否，不存在就记录
		zap.L().Info("用户不存在，请重新操作")
		responseFail(c, 1, "用户不存在，请重新操作", nil)
		return
	}
	//2.获取video_id并判断id存在与否
	video_id := c.Query("video_id")
	vid, err := strconv.Atoi(video_id)
	if err != nil {
		zap.L().Info("视频id不符合固定，请重新操作")
		responseFail(c, 1, "操作的视频有误", nil)
		return
	}
	action_type := c.Query("action_type")
	//3.根据action_type来操作
	if action_type == "1" {
		//添加评论
		ctext := c.Query("comment_text")
		if len(ctext) >= 50 {
			zap.L().Info("评论过长，评论失败")
			responseFail(c, 1, "评论过长，请稍后重试", nil)
			return
		}
		rsc, err := mysql.CommentAction(username, ctext, vid)
		if err != nil {
			responseFail(c, 1, "评论有误，请稍后重试", nil)
			return
		}
		responseFail(c, 0, "评论成功", rsc)
		return

	} else if action_type == "2" {
		//删除评论
		comment_id := c.Query("comment_id")
		cid, err := strconv.Atoi(comment_id)
		if err != nil {
			responseFail(c, 1, "评论失败", nil)
			return
		}
		//判断评论的id存在与否
		if !mysql.CommentIsExist(cid) {
			responseFail(c, 1, "评论不存在", nil)
			return
		}
		//存在就删除评论
		rsc, err := mysql.DeleteAction(username, vid, cid)
		if err != nil {
			responseFail(c, 1, "删除评论失败", nil)
			return
		}
		responseFail(c, 0, "评论成功", rsc)
	} else {
		zap.L().Info("评论操作不符合固定，请重新操作")
		responseFail(c, 1, "评论有误，请稍后重试", nil)
		return
	}
}

func CommentList(c *gin.Context) {
	video := c.Query("video_id")
	videoId, _ := strconv.ParseInt(video, 10, 64)
	myc, err := jwt.ParseToken(c.Query("token")) //解析token，失败就返回并记录错误
	if err != nil {
		zap.L().Info("token有误！错误为:", zap.Error(err))
		responseFail(c, 1, "获取信息失败，请稍后重试", nil)
		return
	}
	userId := myc.UserID
	commentList := mysql.GetCommentListAndUserList(uint64(videoId), uint64(userId))
	zap.L().Info(commentList[0].Content)
	c.JSON(http.StatusOK, gin.H{
		"status_code":  0,
		"status_msg":   "",
		"comment_list": commentList,
	})
}
