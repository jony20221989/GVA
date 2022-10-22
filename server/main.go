package main

import (
	"fmt"
	"go.uber.org/zap"
	"server/global"
	"server/initialize"
	"time"
)

func main() {
	//初始化配置
	global.VP = initialize.InitViper()
	//初始化日志
	global.LOG = initialize.InitZap()
	//初始化jwt
	initialize.InitJwt()
	//初始化缓存
	global.REDIS = initialize.InitRedis()
	//初始化数据库连接  如果没配置或者没初始化 返回nil
	global.DB = initialize.InitDBConn()
	if global.DB != nil {
		//数据库版本迁移
		//	initialize.AutoMigrate(global.DB)
		db, _ := global.DB.DB()
		// 程序结束前关闭数据库连接
		defer db.Close()
	}
	//从db加载黑名单到缓存
	if global.DB != nil {
		initialize.InitBlackList()
	}
	//获取路由
	Router := initialize.InitRouters()
	//获取端口
	port := fmt.Sprintf(":%d", global.CONFIG.Server.Port)
	//启动服务
	s := initialize.InitServer(port, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.LOG.Info("server run success on ", zap.String("address", port))
	global.LOG.Error(s.ListenAndServe().Error())
}
