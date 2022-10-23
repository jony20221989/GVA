package middleware

import (
	"context"
	"github.com/go-redis/redis/v9"
	"go.uber.org/zap"
	"server/global"
)

func InitRedis() *redis.Client {
	redisCfg := global.CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.LOG.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		global.LOG.Info("redis connect ping response:", zap.String("pong", pong))
	}
	return client
}
