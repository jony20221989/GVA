package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/global"
	"server/model/entity"
	"server/model/request"
	"server/model/response"
	"server/utils"
)

type UserApi struct{}

func (b *UserApi) Login(c *gin.Context) {
	var body request.Login
	err := c.ShouldBindJSON(&body)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(body, utils.LoginVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	u := &entity.SysUser{Username: body.Username, Password: body.Password}
	user, err := userService.Login(u)
	if err != nil {
		global.LOG.Error(err.Error(), zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	if user.Enable != 1 {
		global.LOG.Error("登陆失败! 用户被禁止登录!")
		response.FailWithMessage("用户被禁止登录", c)
		return
	}
	jwtApi := ApiGroupApp.JwtApi
	jwtApi.CreateToken(c, *user)
	return
}
