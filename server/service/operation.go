package service

import (
	"server/global"
	"server/model/entity"
	"server/model/request"
)

type OperationRecordService struct{}

func (operationRecordService *OperationRecordService) CreateSysOperationRecord(sysOperationRecord entity.SysOperationRecord) (err error) {
	err = global.DB.Create(&sysOperationRecord).Error
	return err
}

//@description: 批量删除记录

func (operationRecordService *OperationRecordService) DeleteSysOperationRecordByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]entity.SysOperationRecord{}, "id in (?)", ids.Ids).Error
	return err
}

//@description: 删除操作记录

func (operationRecordService *OperationRecordService) DeleteSysOperationRecord(sysOperationRecord entity.SysOperationRecord) (err error) {
	err = global.DB.Delete(&sysOperationRecord).Error
	return err
}

//@description: 根据id获取单条操作记录

func (operationRecordService *OperationRecordService) GetSysOperationRecord(id uint) (sysOperationRecord entity.SysOperationRecord, err error) {
	err = global.DB.Where("id = ?", id).First(&sysOperationRecord).Error
	return
}

//@description: 分页获取操作记录列表

func (operationRecordService *OperationRecordService) GetSysOperationRecordInfoList(info request.SysOperationRecordSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.PageNum - 1)
	// 创建db
	db := global.DB.Model(&entity.SysOperationRecord{})
	var sysOperationRecords []entity.SysOperationRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Method != "" {
		db = db.Where("method = ?", info.Method)
	}
	if info.Path != "" {
		db = db.Where("path LIKE ?", "%"+info.Path+"%")
	}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("id desc").Limit(limit).Offset(offset).Preload("User").Find(&sysOperationRecords).Error
	return sysOperationRecords, total, err
}
