package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"server/global"
	"server/model/response"
	"server/utils"
)

func (b *BaseApi) Debug(c *gin.Context) {
	response.OkWithMessage("测试接口Debug", c)
}

// 当开启多服务器部署时，替换下面的配置，使用redis共享存储验证码
var store = utils.NewDefaultRedisStore()

//var store = base64Captcha.DefaultMemStore

type BaseApi struct{}

// Captcha
// @Tags      Base
// @Summary   生成验证码
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.response{data=systemRes.SysCaptchaResponse,msg=string}  "生成验证码,返回包括随机数id,base64,验证码长度"
// @Router    /base/captcha [post]
func (b *BaseApi) Captcha(c *gin.Context) {
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(global.CONFIG.Captcha.ImgHeight, global.CONFIG.Captcha.ImgWidth, global.CONFIG.Captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(c)) // v9下使用redis
	//	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
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