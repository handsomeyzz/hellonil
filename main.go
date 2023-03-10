package main

import (
	"fmt"
	"hellonil/dao/feed"
	"hellonil/dao/mysql"
	"hellonil/dao/redis"
	"hellonil/logger"
	"hellonil/pkg/jwt"
	"hellonil/pkg/snowflake"
	"hellonil/router"
	"hellonil/setting"
	"os"
)

func Init() {
	if len(os.Args) < 2 {
		fmt.Println("need config file.eg: hellonil config.yaml")
		return
	}
	// 加载配置
	if err := setting.Init(os.Args[1]); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close() // 程序退出关闭数据库连接
	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redis.Close()

	if err := feed.Init(setting.Conf.CosConfig); err != nil {
		fmt.Printf("init cos filed,err:%v\n", err)
		return
	}

	if err := snowflake.Init(setting.Conf.StartTime, setting.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}
	jwt.Init(setting.Conf.JwtConfig)
	// 注册路由
	r := router.SetupRouter(setting.Conf.Mode)
	err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}

func main() {
	Init()
}
