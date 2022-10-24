package request

import "server/model/entity"

type SearchApiParams struct {
	entity.SysApi
	PageInfo
	OrderKey string `json:"orderKey"` // 排序
	Desc     bool   `json:"desc"`     // 排序方式:升序false(默认)|降序true
}
