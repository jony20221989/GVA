package router

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	//userRouter := Router.Group("user") //.Use(middleware.OperationRecord())
	//userRouterWithoutRecord := Router.Group("user")
	//baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	//{
	//	userRouter.POST("admin_register", baseApi.Register) // 管理员注册账号
	//
	//}

}
