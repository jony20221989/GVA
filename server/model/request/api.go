package request

import "server/model/entity"

type SearchApiParams struct {
	entity.SysApi
	PageInfo
	OrderKey string `json:"orderKey"` // 排序

}
