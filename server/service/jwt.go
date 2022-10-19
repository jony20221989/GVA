package service

import (
	"context"
	"go.uber.org/zap"
	"server/global"
	"server/model/entity"
	"server/utils"
)

type JwtService struct{}

func (jwtService *JwtService) JsonInBlacklist(jwtList entity.JwtBlacklist) (err error) {
	err = global.DB.Create(&jwtList).Error
	if err != nil {
		return
	}
	//global.BlackCache.SetDefault(jwtList.Jwt, struct{}{})
	err = global.REDIS.Set(context.Background(), jwtList.Jwt, struct{}{}, 0).Err()
	if err != nil {
		global.LOG.Error("RedisJwtBlacklistSetError!", zap.Error(err))
	}
	return
}

// InBlacklist 判断JWT是否在黑名单内部
func (jwtService *JwtService) InBlacklist(jwt string) bool {
	//_, ok := global.BlackCache.Get(jwt)
	err := global.REDIS.Get(context.Background(), jwt).Err()
	if err != nil {
		return true
	}
	return false
	// err := global.GVA_DB.Where("jwt = ?", jwt).First(&system.JwtBlacklist{}).Error
	// isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	// return !isNotFound
}

func (jwtService *JwtService) GetRedisJWT(userName string) (redisJWT string, err error) {
	redisJWT, err = global.REDIS.Get(context.Background(), userName).Result()
	return redisJWT, err
}

// SetRedisJWT jwt存入redis并设置过期时间
func (jwtService *JwtService) SetRedisJWT(userName string, jwt string) (err error) {
	// 此处过期时间等于jwt过期时间
	dr, err := utils.ParseDuration(global.CONFIG.JWT.ExpiresTime)
	if err != nil {
		return err
	}
	timer := dr
	err = global.REDIS.Set(context.Background(), userName, jwt, timer).Err()
	return err
}
