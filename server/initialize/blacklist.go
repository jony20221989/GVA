package initialize

import (
	"go.uber.org/zap"
	"server/global"
	"server/model/entity"
	"server/utils"
)

func InitBlackList() {
	var data []string
	err := global.DB.Model(&entity.JwtBlacklist{}).Select("jwt").Find(&data).Error
	if err != nil {
		global.LOG.Error("加载数据库jwt黑名单失败!", zap.Error(err))
		return
	}

	for i := 0; i < len(data); i++ {
		err := utils.RedisSetBlank(data[i])
		if err != nil {
			return
		}
	} // jwt黑名单 加入 Redis  中

}
