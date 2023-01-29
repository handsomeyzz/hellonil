package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"hellonil/logger"
	"net/http"
)

func main() {
	err := logger.Init()
	if err != nil {
		zap.L().Debug("初始化出错！", zap.Any("err", err))
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, "666")
	})
	r.Run("127.0.0.1:8080")
}
