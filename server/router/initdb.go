package router

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
)

type InitRouter struct{}

func (s *InitRouter) InitInitRouter(Router *gin.RouterGroup) {
	initRouter := Router.Group("init")
	dbApi := v1.ApiGroupApp.DBApi
	{
		initRouter.POST("initdb", dbApi.InitDB)   // 创建Api
		initRouter.POST("checkdb", dbApi.CheckDB) // 创建Api
	}
}
