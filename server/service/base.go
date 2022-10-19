package service

import (
	"context"
	"github.com/mojocn/base64Captcha"
	"server/global"
	"server/utils"
)

// 当开启多服务器部署时，替换下面的配置，使用redis共享存储验证码
var store = utils.NewDefaultRedisStore()

//var store = base64Captcha.DefaultMemStore

type BaseService struct{}

// CreateCaptcha 创建验证码
func (baseService *BaseService) CreateCaptcha() (id, b64s string, err error) {
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(global.CONFIG.Captcha.ImgHeight, global.CONFIG.Captcha.ImgWidth, global.CONFIG.Captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(context.Background())) // v9下使用redis
	//	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err = cp.Generate()
	return id, b64s, err
}
