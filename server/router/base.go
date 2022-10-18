package router

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	baseApi := v1.ApiGroupApp.BaseApi
	{
		baseRouter.GET("test", baseApi.Debug)
		baseRouter.POST("captcha", baseApi.Captcha)
		baseRouter.POST("login", baseApi.Login)
	}
	return baseRouter
}
