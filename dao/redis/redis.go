package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"hellonil/setting"
)

// 实际生产环境下 context.Background() 按需替换

var (
	client *redis.Client
	ctx    = context.Background()
)

// Init 初始化连接
func Init(cfg *setting.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password:     cfg.Password, // no password set
		DB:           cfg.DB,       // use default DB
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})

	_, err = client.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	_ = client.Close()
}
