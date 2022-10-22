package middleware

import (
	"server/global"
	"server/model/response"
	"server/service"
	"server/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var jwtService = service.ServiceGroupApp.JwtService

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("x-token")
		//token不存在
		if token == "" {
			response.FailWithForbidden("未登录或非法访问", c)
			c.Abort()
			return
		}
		//token失效
		if jwtService.CheckBlacklist(token) {
			response.FailWithForbidden("您的帐户异地登陆或令牌失效", c)
			c.Abort()
			return
		}
		jwt := utils.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := jwt.ParseToken(token)
		if err != nil {
			if err == utils.TokenExpired {
				response.FailWithForbidden("授权已过期", c)
				c.Abort()
				return
			}
			response.FailWithForbidden(err.Error(), c)
			c.Abort()
			return
		}

		// 已登录用户被管理员禁用 需要使该用户的jwt失效 此处比较消耗性能 如果需要 请自行打开
		// 用户被删除的逻辑 需要优化 此处比较消耗性能 如果需要 请自行打开

		//if user, err := userService.FindUserByUuid(claims.UUID.String()); err != nil || user.Enable == 2 {
		//	_ = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: token})
		//	response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
		//	c.Abort()
		//}

		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			dr, _ := utils.ParseDuration(global.CONFIG.JWT.ExpiresTime)
			claims.ExpiresAt = time.Now().Add(dr).Unix()
			newToken, _ := jwt.CreateTokenByOldToken(token, *claims)
			newClaims, _ := jwt.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
			if global.CONFIG.Server.UseMultipoint {
				RedisJwtToken, err := utils.RedisGetJWT(newClaims.Username)
				if err != nil {
					global.LOG.Error("get redis jwt failed", zap.Error(err))
				} else {
					// 当之前的取成功时才进行拉黑操作
					_ = utils.RedisSetBlank(RedisJwtToken)
				}
				// 无论如何都要记录当前的活跃状态
				_ = utils.RedisSetJWT(newClaims.Username, newToken)
			}
		}
		c.Set("claims", claims)
		c.Next()
	}
}
