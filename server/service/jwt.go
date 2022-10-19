package service

import (
	"server/global"
	"server/model/entity"
	"server/utils"
)

type JwtService struct{}

// SetBlacklist jwtList设置到黑名单里
func (jwtService *JwtService) SetBlacklist(jwtList entity.JwtBlacklist) (err error) {
	err = global.DB.Create(&jwtList).Error
	if err != nil {
		return
	}
	//global.BlackCache.SetDefault(jwtList.Jwt, struct{}{})
	//err = global.REDIS.Set(context.Background(), jwtList.Jwt, struct{}{}, 0).Err()
	//if err != nil {
	//	global.LOG.Error("RedisJwtBlacklistSetError!", zap.Error(err))
	//}
	//return err
	err = utils.RedisSetBlank(jwtList.Jwt)
	return
}

// CheckBlacklist 判断JWT是否在黑名单内部
func (jwtService *JwtService) CheckBlacklist(jwt string) bool {
	b, err := utils.RedisCheckBlank(jwt)
	//err := global.REDIS.Get(context.Background(), jwt).Err()
	if err != nil {
		return b
	}
	return false
	// err := global.GVA_DB.Where("jwt = ?", jwt).First(&system.JwtBlacklist{}).Error
	// isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	// return !isNotFound
}
