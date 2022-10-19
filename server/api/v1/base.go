package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/global"
	"server/model/response"
)

type BaseApi struct{}

func (b *BaseApi) Debug(c *gin.Context) {
	response.OkWithMessage("测试接口Debug", c)
}

// Captcha 生成验证码api
func (b *BaseApi) Captcha(c *gin.Context) {
	id, b64s, err := baseService.CreateCaptcha()
	if err != nil {
		global.LOG.Error("验证码获取失败!", zap.Error(err))
		response.FailWithMessage("验证码获取失败", c)
		return
	}
	response.OkWithDetailed(response.CaptchaResponse{
		CaptchaId:     id,
		PicPath:       b64s,
		CaptchaLength: global.CONFIG.Captcha.KeyLong,
	}, "验证码获取成功", c)
}
