package router

import (
	"hellonil/controller"
	"hellonil/logger"
	"hellonil/middlewares"
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	r := gin.New()
	//r.Use(logger.GinLogger(), logger.GinRecovery(true), middlewares.RateLimitMiddleware(2*time.Second, 1))
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK!")
	})
	//r.Group("/douyin")
	//基本接口
	r.GET("/douyin/feed/", controller.FeedSearch)         //feed流接口
	r.POST("/douyin/user/register/", controller.Register) //用户注册
	r.POST("/douyin/user/login/", controller.Login)       //用户登录
	r.GET("/douyin/user/", controller.UserMsg)
	r.POST("/douyin/publish/action/", controller.PublishAction) //投稿
	r.GET("/douyin/publish/list/", controller.PublishList)      //发布列表
	//互动接口
	r.POST("/douyin/favorite/action/", controller.Approve)      //点赞  //这儿可以添加中间件
	r.POST("/douyin/comment/action/", controller.CommentAction) //评论操作
	r.GET("/douyin/comment/list/", controller.CommentList)      //评论列表
	r.GET("/douyin/favorite/list/", controller.FavoriteList)    //喜爱列表

	//社交接口
	r.POST("/douyin/relation/action/", middlewares.JWTAuth(), controller.FollowAction) //关注操作
	r.GET("/douyin/relation/follower/list/", controller.FollowList)                    //关注列表
	r.GET("/douyin/douyin/relation/follower/list/", controller.FollowerList)           //粉丝列表
	r.POST("/douyin/message/action/", middlewares.JWTAuth(), controller.SendMsg)       //发送消息

	pprof.Register(r) // 注册pprof相关路由

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
