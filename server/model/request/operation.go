package request

import (
	"server/model/entity"
)

type SysOperationRecordSearch struct {
	entity.SysOperationRecord
	PageInfo
}
