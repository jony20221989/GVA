package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"go.uber.org/zap"
	"server/global"
	"server/model/entity"
	"server/model/request"
	"server/model/response"
	"server/utils"
)

type JwtApi struct{}

// CreateToken 登录以后签发jwt
func (b *JwtApi) CreateToken(c *gin.Context, user entity.SysUser) {
	jwt := &utils.JWT{SigningKey: []byte(global.CONFIG.JWT.SigningKey)} // 唯一签名
	claims := jwt.CreateClaims(request.BaseClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
	})
	token, err := jwt.CreateToken(claims)
	if err != nil {
		global.LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	if !global.CONFIG.Server.UseMultipoint {
		response.OkWithDetailed(response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
		return
	}
	fmt.Println("token:" + token)

	redisJWT, err := utils.RedisGetJWT(user.Username) //jwtService.GetRedisJWT(user.Username)
	//key 不存在
	if err == redis.Nil {
		//在redis中保存token
		err := utils.RedisSetJWT(user.Username, token)
		//保存失败
		if err != nil {
			global.LOG.Error("登录保存token失败", zap.Error(err))
			response.FailWithMessage("登录保存token失败", c)
			return
		}
		//保存成功
		response.OkWithDetailed(response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
		return
	}
	//获取redis jwt失败
	if err != nil {
		global.LOG.Error("登录获取redis jwt失败")
		response.FailWithDetailed(err.Error(), "登录获取redis jwt失败", c)
		return
	}
	//获取redis 之前的jwt成功
	var blackJWT entity.JwtBlacklist
	blackJWT.Jwt = redisJWT
	//作废之前的jwt
	if err := jwtService.SetBlacklist(blackJWT); err != nil {
		response.FailWithMessage("之前的jwt作废失败", c)
		return
	}
	if err := utils.RedisSetJWT(user.Username, token); err != nil {
		global.LOG.Error("登录保存token失败", zap.Error(err))
		response.FailWithMessage("登录保存token失败", c)
		return
	}
	response.OkWithDetailed(response.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}, "登录成功", c)

}
