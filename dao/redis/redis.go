package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"hellonil/config"
)

var ctx = context.Background()
var Re *redis.Client

func Init() error {
	Re = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.RedisX().Addr, config.RedisX().Port),
		Password: fmt.Sprintf("%s", config.RedisX().Password), // no password set
		DB:       0,                                           // use default DB
	})
	_, err := Re.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	_ = Re.Close()
}
