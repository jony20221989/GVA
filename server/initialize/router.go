package initialize

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"server/global"
	"server/router"
)

// 初始化总路由

func InitRouters() *gin.Engine {
	engine := gin.Default()
	routerGroup := router.RouterGroupApp
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
		routerGroup.InitBaseRouter(PublicGroup)
	}

	/*PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{

	}*/
	global.LOG.Info("router register success")
	return engine
}
