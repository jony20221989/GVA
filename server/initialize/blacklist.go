package initialize

import (
	"context"
	"go.uber.org/zap"
	"server/config"
	"server/global"
	"server/model/entity"
)

func InitBlackList() {
	var data []string
	err := global.DB.Model(&entity.JwtBlacklist{}).Select("jwt").Find(&data).Error
	if err != nil {
		global.LOG.Error("加载数据库jwt黑名单失败!", zap.Error(err))
		return
	}

	for i := 0; i < len(data); i++ {
		global.REDIS.SAdd(context.Background(), config.BLACKLIST_KEY, data[i])
	} // jwt黑名单 加入 Redis  中

}
