package middleware

import (
	"server/global"
	"server/utils"
)

func InitJwt() {
	dr, err := utils.ParseDuration(global.CONFIG.JWT.ExpiresTime)
	if err != nil {
		panic(err)
	}
	_, err = utils.ParseDuration(global.CONFIG.JWT.BufferTime)
	if err != nil {
		panic(err)
	}
	global.LOG.Info(dr)
}
