package initialize

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"server/global"
	"server/middleware"
	"server/router"
)

// 初始化总路由

func InitRouters() *gin.Engine {
	engine := gin.Default()
	routerGroupApp := router.RouterGroupApp
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.LOG.Info("register swagger handler")

	//公共路由
	PublicGroup := engine.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	{
		// 注册基础功能路由 不做鉴权
		routerGroupApp.InitBaseRouter(PublicGroup)
		// 数据库初始化路由
		routerGroupApp.InitDBRouter(PublicGroup)
	}

	PrivateGroup := engine.Group("")
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		routerGroupApp.InitApiRouter(PrivateGroup)
		routerGroupApp.InitUserRouter(PrivateGroup)
	}
	global.LOG.Info("router register success")
	return engine
}
