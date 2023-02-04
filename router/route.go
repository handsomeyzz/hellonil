package router

import (
	"hellonil/controller"
	"hellonil/logger"
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
	r.GET("/douyin/feed/", controller.FeedSearch)
	r.POST("/douyin/user/register/", controller.Register)       //用户注册
	r.POST("/douyin/user/login/", controller.Login)             //用户登录
	r.POST("/douyin/publish/action/", controller.PublishAction) //投稿
	r.POST("/favorite/action/", controller.Approve)             //点赞
	r.GET("/user/", controller.UserMsg)

	pprof.Register(r) // 注册pprof相关路由

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
