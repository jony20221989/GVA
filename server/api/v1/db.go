package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/global"
	"server/model/request"
	"server/model/response"
)

type DBApi struct{}

func (i *DBApi) Init(c *gin.Context) {
	if global.DB != nil {
		global.LOG.Error("已存在数据库配置!")
		response.FailWithMessage("已存在数据库配置", c)
		return
	}
	var dbInfo request.InitDB
	if err := c.ShouldBindJSON(&dbInfo); err != nil {
		global.LOG.Error("参数校验不通过!", zap.Error(err))
		response.FailWithMessage("参数校验不通过", c)
		return
	}
	if err := initDBService.InitDB(dbInfo); err != nil {
		global.LOG.Error("自动创建数据库失败!", zap.Error(err))
		response.FailWithMessage("自动创建数据库失败，请查看后台日志，检查后在进行初始化", c)
		return
	}
	response.OkWithMessage("自动创建数据库成功", c)
}

func (i *DBApi) Check(c *gin.Context) {
	var (
		message  = "前往初始化数据库"
		needInit = true
	)

	if global.DB != nil {
		message = "数据库无需初始化"
		needInit = false
	}
	global.LOG.Info(message)
	response.OkWithDetailed(gin.H{"needInit": needInit}, message, c)
}
