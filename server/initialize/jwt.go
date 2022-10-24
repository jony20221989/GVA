package initialize

import (
	"server/global"
	"server/utils"
)

func InitJwt() {
	dr, err := utils.ParseDuration(global.CONFIG.JWT.ExpiresTime)
	if err != nil {
		panic(err)
	}
	bf, err := utils.ParseDuration(global.CONFIG.JWT.BufferTime)
	if err != nil {
		panic(err)
	}
	global.LOG.Info("初始化dr:", dr)
	global.LOG.Info("初始化bf:", bf)
}
