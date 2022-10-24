package router

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	//userRouter := Router.Group("user") //.Use(middleware.OperationRecord())
	userRouterWithoutRecord := Router.Group("user")
	userApi := v1.ApiGroupApp.UserApi
	{
		userRouterWithoutRecord.GET("getUserInfo", userApi.GetUserInfo)
		userRouterWithoutRecord.POST("getUserList", userApi.GetUserList)

	}

}
