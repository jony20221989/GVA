package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/global"
	"server/model/request"
	"server/model/response"
	"server/utils"
)

type SystemApiApi struct{}

func (s *SystemApiApi) GetAllApis(c *gin.Context) {
	apis, err := apiService.GetAllApis()
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.SysAPIListResponse{Apis: apis}, "获取成功", c)
}

func (s *SystemApiApi) GetApiList(c *gin.Context) {
	var pageInfo request.SearchApiParams
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := apiService.GetAPIInfoList(pageInfo.SysApi, pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc)
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithData(response.PageResult{
		List:     list,
		Total:    total,
		PageNum:  pageInfo.PageNum,
		PageSize: pageInfo.PageSize,
	}, c)
}
