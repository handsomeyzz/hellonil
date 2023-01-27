package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hellonil/dao/feed"
	"io"
	"net/http"
)

//	func main() {
//		if err := mysql.Init(); err != nil {
//			fmt.Println("mysql初始化失败！")
//		}
//		if err := redis.Init(); err != nil {
//			fmt.Println("redis初始化失败")
//			fmt.Println(err)
//		}
//
// }
func main() {
	if err := feed.Init(); err != nil {
		fmt.Println(err)
	}
	r := gin.Default()
	r.POST("/hello", func(c *gin.Context) {
		file, err := c.FormFile("data")
		if err != nil {
			c.String(http.StatusOK, "请求失败")
		}
		fileName := file.Filename
		fmt.Println(fileName)
		var r io.Reader
		r, _ = file.Open()
		url, err := feed.UploadVideo(fileName, r)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(url)
	})
	r.Run("127.0.0.1:8088")
}
