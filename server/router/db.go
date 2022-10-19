package router

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
)

type InitRouter struct{}

func (s *InitRouter) InitDBRouter(Router *gin.RouterGroup) {
	initRouter := Router.Group("db")
	dbApi := v1.ApiGroupApp.DBApi
	{
		initRouter.POST("check", dbApi.Check) // 检查是否已经初始化过数据库
		initRouter.POST("init", dbApi.Init)   // 初始化数据库
	}
}
