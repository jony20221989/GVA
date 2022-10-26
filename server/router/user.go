package router

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
	"server/middleware"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user").Use(middleware.OperationRecord())
	userRouterWithoutRecord := Router.Group("user")
	userApi := v1.ApiGroupApp.UserApi
	{
		userRouter.POST("admin_register", userApi.Register)               // 管理员注册账号
		userRouter.POST("changePassword", userApi.ChangePassword)         // 用户修改密码
		userRouter.POST("setUserAuthority", userApi.SetUserAuthority)     // 设置用户权限
		userRouter.DELETE("deleteUser", userApi.DeleteUser)               // 删除用户
		userRouter.PUT("setUserInfo", userApi.SetUserInfo)                // 设置用户信息
		userRouter.PUT("setSelfInfo", userApi.SetSelfInfo)                // 设置自身信息
		userRouter.POST("setUserAuthorities", userApi.SetUserAuthorities) // 设置用户权限组
		userRouter.POST("resetPassword", userApi.ResetPassword)           // 设置用户权限组
	}
	{
		userRouterWithoutRecord.GET("getUserInfo", userApi.GetUserInfo)
		userRouterWithoutRecord.POST("getUserList", userApi.GetUserList)

	}

}
