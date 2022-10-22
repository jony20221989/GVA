package utils

import (
	"context"
	"go.uber.org/zap"
	"server/config"
	"server/global"
	"time"
)

func RedisSetBlank(jwt string) (err error) {

	err = global.REDIS.SAdd(context.Background(), config.BLACKLIST_KEY, jwt).Err()
	if err != nil {
		global.LOG.Error("Redis 黑名单设置错误", zap.Error(err))
	}
	return err
}

func RedisGetBlank() (members []string, err error) {
	members, err = global.REDIS.SMembers(context.Background(), config.BLACKLIST_KEY).Result()
	if err != nil {
		global.LOG.Error("Redis 获取黑名单错误", zap.Error(err))
	}
	return
}

func RedisCheckBlank(jwt string) (b bool, err error) {
	b, err = global.REDIS.SIsMember(context.Background(), config.BLACKLIST_KEY, jwt).Result()
	if err != nil {
		global.LOG.Error("Redis 获取黑名单错误", zap.Error(err))
	}
	return
}

func RedisGetJWT(userName string) (redisJWT string, err error) {
	//redisJWT, err = global.REDIS.Get(context.Background(), userName).Result()
	redisJWT, err = global.REDIS.Get(context.Background(), config.ONLINE_KEY+userName).Result()
	return
}

func RedisSetJWT(userName string, jwt string) (err error) {
	// 此处过期时间等于jwt过期时间
	dr, err := ParseDuration(global.CONFIG.JWT.ExpiresTime)
	if err != nil {
		return err
	}
	timer := dr
	err = global.REDIS.Set(context.Background(), config.ONLINE_KEY+userName, jwt, timer*time.Second).Err()
	return err

}
